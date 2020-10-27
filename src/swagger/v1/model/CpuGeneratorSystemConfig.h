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
 * CpuGeneratorSystemConfig.h
 *
 * System CPU load
 */

#ifndef CpuGeneratorSystemConfig_H_
#define CpuGeneratorSystemConfig_H_


#include "ModelBase.h"


namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// System CPU load
/// </summary>
class  CpuGeneratorSystemConfig
    : public ModelBase
{
public:
    CpuGeneratorSystemConfig();
    virtual ~CpuGeneratorSystemConfig();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// CpuGeneratorSystemConfig members

    /// <summary>
    /// CPU utilization value. The maximum value is (100 * core_count). 
    /// </summary>
    double getUtilization() const;
    void setUtilization(double value);
    
protected:
    double m_Utilization;

};

}
}
}

#endif /* CpuGeneratorSystemConfig_H_ */