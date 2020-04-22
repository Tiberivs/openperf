#ifndef _OP_PACKET_GENERATOR_TRAFFIC_HEADER_UTILS_HPP_
#define _OP_PACKET_GENERATOR_TRAFFIC_HEADER_UTILS_HPP_

#include "packet/generator/traffic/header/config.hpp"
#include "packet/generator/traffic/header/container.hpp"

namespace openperf::packet::generator::traffic::header {

using config_index_type = decltype(std::declval<config_instance>().index());
using config_key = std::vector<config_index_type>;

config_container update_context_fields(config_container&&) noexcept;

size_t count_headers(const config_container&, modifier_mux) noexcept;

container make_headers(const config_container&, modifier_mux) noexcept;

config_key get_config_key(const config_container&) noexcept;

void update_lengths(const config_key& indexes,
                    uint8_t packet[],
                    uint16_t pkt_length) noexcept;

} // namespace openperf::packet::generator::traffic::header

#endif /* _OP_PACKET_GENERATOR_TRAFFIC_HEADER_UTILS_HPP_ */