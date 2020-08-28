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
 * GetPacketCapturesPcapConfig.h
 *
 * Parameters for the capture data retrieval
 */

#ifndef GetPacketCapturesPcapConfig_H_
#define GetPacketCapturesPcapConfig_H_


#include "ModelBase.h"

#include <string>
#include <vector>

namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// Parameters for the capture data retrieval
/// </summary>
class  GetPacketCapturesPcapConfig
    : public ModelBase
{
public:
    GetPacketCapturesPcapConfig();
    virtual ~GetPacketCapturesPcapConfig();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// GetPacketCapturesPcapConfig members

    /// <summary>
    /// List of capture results identifiers
    /// </summary>
    std::vector<std::string>& getIds();
        /// <summary>
    /// The packet offset in the capture buffer to start reading (0 based)
    /// </summary>
    int64_t getPacketStart() const;
    void setPacketStart(int64_t value);
    bool packetStartIsSet() const;
    void unsetPacket_start();
    /// <summary>
    /// The packet offset in the capture buffer to end reading (0 based)
    /// </summary>
    int64_t getPacketEnd() const;
    void setPacketEnd(int64_t value);
    bool packetEndIsSet() const;
    void unsetPacket_end();

protected:
    std::vector<std::string> m_Ids;

    int64_t m_Packet_start;
    bool m_Packet_startIsSet;
    int64_t m_Packet_end;
    bool m_Packet_endIsSet;
};

}
}
}

#endif /* GetPacketCapturesPcapConfig_H_ */