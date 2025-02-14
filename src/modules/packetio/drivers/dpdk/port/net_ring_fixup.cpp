#include <cstring>

#include "packetio/drivers/dpdk/mbuf_metadata.hpp"
#include "packetio/drivers/dpdk/names.hpp"
#include "packetio/drivers/dpdk/port/net_ring_fixup.hpp"
#include "packetio/drivers/dpdk/port_info.hpp"

namespace openperf::packetio::dpdk::port {

std::string callback_net_ring_fixup::description()
{
    return ("net_ring metadata fixup");
}

static uint16_t fixup_net_ring_metadata([[maybe_unused]] uint16_t port_id,
                                        [[maybe_unused]] uint16_t queue_id,
                                        rte_mbuf* packets[],
                                        uint16_t nb_packets,
                                        [[maybe_unused]] uint16_t max_packets,
                                        [[maybe_unused]] void* user_param)
{
    std::for_each(packets, packets + nb_packets, mbuf_tx_sink_clear);
    return (nb_packets);
}

rx_callback<callback_net_ring_fixup>::rx_callback_fn
callback_net_ring_fixup::callback()
{
    return (fixup_net_ring_metadata);
}

static net_ring_fixup::variant_type make_net_ring_fixup(uint16_t port_id)
{
    if (port_info::driver_name(port_id) == driver_names::ring) {
        return (callback_net_ring_fixup(port_id));
    }

    return (null_feature(port_id));
}

net_ring_fixup::net_ring_fixup(uint16_t port_id)
    : feature(make_net_ring_fixup(port_id))
{}

} // namespace openperf::packetio::dpdk::port
