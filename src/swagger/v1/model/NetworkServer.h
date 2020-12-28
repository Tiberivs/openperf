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
 * NetworkServer.h
 *
 * Network server
 */

#ifndef NetworkServer_H_
#define NetworkServer_H_


#include "ModelBase.h"

#include "NetworkServerStats.h"
#include <string>

namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// Network server
/// </summary>
class  NetworkServer
    : public ModelBase
{
public:
    NetworkServer();
    virtual ~NetworkServer();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// NetworkServer members

    /// <summary>
    /// Unique network server identifier
    /// </summary>
    std::string getId() const;
    void setId(std::string value);
        /// <summary>
    /// UDP/TCP port on which to listen for client connections
    /// </summary>
    int32_t getPort() const;
    void setPort(int32_t value);
        /// <summary>
    /// Protocol which is used to establish client connections
    /// </summary>
    std::string getProtocol() const;
    void setProtocol(std::string value);
        /// <summary>
    /// Bind server socket to a particular device, specified as interface name (required for dpdk driver)
    /// </summary>
    std::string getInterface() const;
    void setInterface(std::string value);
    bool interfaceIsSet() const;
    void unsetInterface();
    /// <summary>
    /// 
    /// </summary>
    std::shared_ptr<NetworkServerStats> getStats() const;
    void setStats(std::shared_ptr<NetworkServerStats> value);
    
protected:
    std::string m_Id;

    int32_t m_Port;

    std::string m_Protocol;

    std::string m_Interface;
    bool m_InterfaceIsSet;
    std::shared_ptr<NetworkServerStats> m_Stats;

};

}
}
}

#endif /* NetworkServer_H_ */