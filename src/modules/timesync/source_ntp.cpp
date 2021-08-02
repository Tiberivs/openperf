#include <cinttypes>
#include <string>

#include <arpa/inet.h>

#include "core/op_log.h"
#include "core/op_event_loop.hpp"
#include "timesync/api.hpp"
#include "timesync/clock.hpp"
#include "timesync/ntp.hpp"
#include "timesync/source_ntp.hpp"

namespace openperf::timesync::source {

static constexpr auto poll_period_adjust = 4s;

static int handle_ntp_poll(const struct op_event_data* data, void* arg);

ntp::ntp(core::event_loop& loop_, class clock* clock_, struct addrinfo* ai_)
    : loop(loop_)
    , addrinfo(ai_)
    , socket(ai_)
    , clock(clock_)
{
    auto callbacks = op_event_callbacks{.on_timeout = handle_ntp_poll};
    loop.get().add((100ns).count(), &callbacks, this, &poll_loop_id);
    assert(poll_loop_id);
}

ntp::~ntp()
{
    assert(poll_loop_id);
    loop.get().del(poll_loop_id);

    assert(socket.fd() != -1);
    loop.get().del(socket.fd());
}

static void* get_address(const struct addrinfo* ai)
{
    return (ai->ai_family == AF_INET
                ? static_cast<void*>(std::addressof(
                    reinterpret_cast<sockaddr_in*>(ai->ai_addr)->sin_addr))
                : static_cast<void*>(std::addressof(
                    reinterpret_cast<sockaddr_in6*>(ai->ai_addr)->sin6_addr)));
}

static in_port_t get_port(const struct addrinfo* ai)
{
    return (
        ntohs(ai->ai_family == AF_INET
                  ? reinterpret_cast<sockaddr_in*>(ai->ai_addr)->sin_port
                  : reinterpret_cast<sockaddr_in6*>(ai->ai_addr)->sin6_port));
}

std::pair<std::string, std::string> get_name_and_port(const struct addrinfo* ai)
{
    auto name = std::string{};
    auto port = std::string{};

    name.resize(api::name_max_length);
    port.resize(api::port_max_length);

    if (getnameinfo(ai->ai_addr,
                    ai->ai_addrlen,
                    name.data(),
                    name.size(),
                    port.data(),
                    port.size(),
                    NI_NAMEREQD)
        == 0) {
        name.resize(strnlen(name.data(), name.capacity()));
        port.resize(strnlen(port.data(), port.capacity()));

        return (std::make_pair(name, port));
    }

    /* Resolution failed for some reason; return the raw socket data */
    static_assert(api::name_max_length > INET6_ADDRSTRLEN);

    if (inet_ntop(
            ai->ai_family, get_address(ai), name.data(), ai->ai_addrlen)) {
        name.resize(strnlen(name.data(), name.capacity()));
        return (std::make_pair(name, std::to_string(get_port(ai))));
    }

    /* All forms of resolution failed */
    return (std::make_pair("unknown", "unknown"));
}

/**
 * Get the poll delay for the i'th packet.
 * This formula gives us a nice exponential back-off for
 * i = [0, ntp_startup_packets] which allows us to quickly get
 * lots of NTP timestamps when we start while slowing our polling
 * rate down afterwards.
 */
std::chrono::nanoseconds ntp_poll_delay(unsigned i,
                                        std::chrono::seconds max_period)
{
    using seconds = std::chrono::duration<double>;
    static constexpr int ntp_startup_packets = 8;

    auto period = seconds(
        i < ntp_startup_packets ? std::pow(
            std::exp(std::log(max_period.count()) / ntp_startup_packets), i)
                                : max_period.count());

    return (std::chrono::duration_cast<std::chrono::nanoseconds>(period));
}

static void handle_invalid_ntp_reply(source::ntp* source,
                                     timesync::ntp::packet& reply)
{
    auto np = get_name_and_port(source->addrinfo.get());
    switch (static_cast<timesync::ntp::kiss_code>(reply.refid)) {
    case timesync::ntp::kiss_code::KISS_DENY:
    case timesync::ntp::kiss_code::KISS_RSTR:
        OP_LOG(OP_LOG_WARNING,
               "NTP server %s has denied access; ceasing to poll\n",
               np.first.c_str());
        source->stats.poll_period =
            std::numeric_limits<std::chrono::seconds>::max();
        break;
    case timesync::ntp::kiss_code::KISS_RATE:
        OP_LOG(OP_LOG_WARNING,
               "Received RATE notification from NTP server %s; "
               "increasing poll period %ld -> %ld\n",
               np.first.c_str(),
               source->stats.poll_period.count(),
               (source->stats.poll_period + poll_period_adjust).count());
        source->stats.poll_period += poll_period_adjust;
        break;
    default: {
        /* refid contains a 4 octet ASCII string in network order; ugh */
        auto refid = ntohl(reply.refid);
        OP_LOG(OP_LOG_WARNING,
               "Received invalid reply from NTP server %s; kiss code = %.4s\n",
               np.first.c_str(),
               reinterpret_cast<const char*>(&refid));
    }
    }
}

static int handle_ntp_reply(const struct op_event_data*, void* arg)
{
    auto* source = reinterpret_cast<source::ntp*>(arg);

    auto buffer = std::array<std::byte, timesync::ntp::packet_size>();
    auto length = buffer.size();
    while (auto rx_time = source->socket.recv(buffer.data(), length)) {
        assert(length <= buffer.size());
        auto reply = timesync::ntp::deserialize(buffer.data(), length);

        /* Skip replies we can't decode. */
        if (!reply) { continue; }

        /*
         * According to RFC 5905, if the stratum is 0, then the
         * reply is invalid and the server might be trying to tell
         * us something.
         */
        if (!reply->stratum) {
            handle_invalid_ntp_reply(source, *reply);
            continue;
        }

        /*
         * Finally, verify that this is a reply to a request we sent
         * by confirming that it contains an origin that matches what
         * we expect.
         */
        if (reply->origin != source->expected_origin) {
            OP_LOG(OP_LOG_WARNING,
                   "Ignoring NTP reply with unrecognized origin timestamp, "
                   "%016" PRIx64 ".%016" PRIx64 "\n",
                   reply->origin.bt_sec,
                   reply->origin.bt_frac);
            continue;
        }

        // ntp::dump(stderr, *reply);

        source->stats.rx++;
        source->stats.stratum = reply->stratum;
        source->clock->update(
            source->last_tx, reply->receive, reply->transmit, *rx_time);
    }

    return (-1);
}

/*
 * Since NTP only uses 64 bits for the timestamp, we need to
 * trim our bintime values accordingly.
 */
static bintime to_ntp_bintime(const bintime& from)
{
    static constexpr int64_t lo_mask = 0xffffffff;
    static constexpr int64_t hi_mask = lo_mask << 32;
    return (bintime{.bt_sec = from.bt_sec & lo_mask,
                    .bt_frac = from.bt_frac & hi_mask});
}

static int handle_ntp_poll(const struct op_event_data* data, void* arg)
{
    auto* source = reinterpret_cast<source::ntp*>(arg);

    auto now = to_bintime(chrono::realtime::now().time_since_epoch());

    /*
     * Record this timestamp so we can verify replies. The server
     * should use our transmit timestamp for the origin timestamp.
     */
    source->expected_origin = to_ntp_bintime(now);

    /* Only need the mode and transmit timestamps for a client request */
    auto request = timesync::ntp::packet{
        .mode = timesync::ntp::mode::MODE_CLIENT, .transmit = now};

    auto buffer = timesync::ntp::serialize(request);

    auto reply_callbacks = op_event_callbacks{.on_read = handle_ntp_reply};
    op_event_loop_add(data->loop, source->socket.fd(), &reply_callbacks, arg);

    auto tx_time = source->socket.send(buffer.data(), buffer.size());
    if (!tx_time) {
        OP_LOG(OP_LOG_ERROR,
               "Failed to send NTP request: %s\n",
               strerror(tx_time.error()));
        op_event_loop_del(data->loop, source->socket.fd());
        return (-1);
    }

    source->last_tx = *tx_time;
    source->stats.tx++;

    op_event_loop_update(
        data->loop,
        source->poll_loop_id,
        static_cast<uint64_t>(
            ntp_poll_delay(source->stats.tx, source->stats.poll_period)
                .count()));

    return (0);
}

api::time_source_ntp to_time_source(const ntp& source)
{
    auto config = api::time_source_config_ntp{};
    auto [name, port] = get_name_and_port(source.addrinfo.get());
    std::copy_n(name.data(), name.size(), config.node);
    std::copy_n(port.data(), port.size(), config.port);

    auto ntp = api::time_source_ntp{
        .config = config,
        .stats = {
            .poll_period = std::chrono::duration_cast<std::chrono::seconds>(
                ntp_poll_delay(source.stats.tx, source.stats.poll_period)),
            .rx_packets = source.stats.rx,
            .tx_packets = source.stats.tx,
            .stratum = source.stats.stratum}};

    return (ntp);
};

} // namespace openperf::timesync::source