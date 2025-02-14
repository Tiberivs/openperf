#ifndef _OP_PACKETIO_MEMORY_DPDK_MEMPOOL_HPP_
#define _OP_PACKETIO_MEMORY_DPDK_MEMPOOL_HPP_

#include <cstdint>
#include <string>

struct rte_mempool;

namespace openperf::packetio::dpdk::mempool {

/*
 * The next two functions are intended for acquiring/releasing memory pools for
 * transmit purposes.
 */
rte_mempool* acquire(unsigned numa_node,
                     uint16_t packet_length,
                     uint16_t packet_count,
                     uint16_t cache_size);

void release(const rte_mempool*);

/* Grab a default memory pool for the specified numa node */
rte_mempool* get_default(unsigned numa_node);

} // namespace openperf::packetio::dpdk::mempool

#endif /* _OP_PACKETIO_MEMORY_DPDK_MEMPOOL_HPP_ */
