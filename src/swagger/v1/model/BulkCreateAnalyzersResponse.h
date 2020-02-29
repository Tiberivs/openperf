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
 * BulkCreateAnalyzersResponse.h
 *
 * 
 */

#ifndef BulkCreateAnalyzersResponse_H_
#define BulkCreateAnalyzersResponse_H_


#include "ModelBase.h"

#include "Analyzer.h"
#include <vector>

namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// 
/// </summary>
class  BulkCreateAnalyzersResponse
    : public ModelBase
{
public:
    BulkCreateAnalyzersResponse();
    virtual ~BulkCreateAnalyzersResponse();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// BulkCreateAnalyzersResponse members

    /// <summary>
    /// List of packet analyzers
    /// </summary>
    std::vector<std::shared_ptr<Analyzer>>& getItems();
    
protected:
    std::vector<std::shared_ptr<Analyzer>> m_Items;

};

}
}
}

#endif /* BulkCreateAnalyzersResponse_H_ */