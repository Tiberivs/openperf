/**
* OpenPerf API
* REST API interface for OpenPerf
*
* OpenAPI spec version: 1
* Contact: support@spirent.com
*
* NOTE: This class is auto generated by the swagger code generator program.
* https://github.com/swagger-api/swagger-codegen.git
* Do not edit the class manually.
*/
/*
 * PacketAnalyzerFlowCounters_sequence.h
 *
 * Advanced Sequencing result counters
 */

#ifndef PacketAnalyzerFlowCounters_sequence_H_
#define PacketAnalyzerFlowCounters_sequence_H_


#include "ModelBase.h"


namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// Advanced Sequencing result counters
/// </summary>
class  PacketAnalyzerFlowCounters_sequence
    : public ModelBase
{
public:
    PacketAnalyzerFlowCounters_sequence();
    virtual ~PacketAnalyzerFlowCounters_sequence();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// PacketAnalyzerFlowCounters_sequence members

    /// <summary>
    /// Number of packets expected but not yet received
    /// </summary>
    int64_t getDropped() const;
    void setDropped(int64_t value);
        /// <summary>
    /// Number of duplicate packets received
    /// </summary>
    int64_t getDuplicate() const;
    void setDuplicate(int64_t value);
        /// <summary>
    /// Number of late packets received
    /// </summary>
    int64_t getLate() const;
    void setLate(int64_t value);
        /// <summary>
    /// Number of reordered packets received
    /// </summary>
    int64_t getReordered() const;
    void setReordered(int64_t value);
        /// <summary>
    /// Number of packets received in the expected sequence
    /// </summary>
    int64_t getInOrder() const;
    void setInOrder(int64_t value);
        /// <summary>
    /// Number of packets received in sequence
    /// </summary>
    int64_t getRunLength() const;
    void setRunLength(int64_t value);
    
protected:
    int64_t m_Dropped;

    int64_t m_Duplicate;

    int64_t m_Late;

    int64_t m_Reordered;

    int64_t m_In_order;

    int64_t m_Run_length;

};

}
}
}

#endif /* PacketAnalyzerFlowCounters_sequence_H_ */