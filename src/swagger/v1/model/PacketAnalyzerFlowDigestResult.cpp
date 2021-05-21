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


#include "PacketAnalyzerFlowDigestResult.h"

namespace swagger {
namespace v1 {
namespace model {

PacketAnalyzerFlowDigestResult::PacketAnalyzerFlowDigestResult()
{
    
}

PacketAnalyzerFlowDigestResult::~PacketAnalyzerFlowDigestResult()
{
}

void PacketAnalyzerFlowDigestResult::validate()
{
    // TODO: implement validation
}

nlohmann::json PacketAnalyzerFlowDigestResult::toJson() const
{
    nlohmann::json val = nlohmann::json::object();

    {
        nlohmann::json jsonArray;
        for( auto& item : m_Centroids )
        {
            jsonArray.push_back(ModelBase::toJson(item));
        }
        val["centroids"] = jsonArray;
            }
    

    return val;
}

void PacketAnalyzerFlowDigestResult::fromJson(nlohmann::json& val)
{
    {
        m_Centroids.clear();
        nlohmann::json jsonArray;
                for( auto& item : val["centroids"] )
        {
            
            if(item.is_null())
            {
                m_Centroids.push_back( std::shared_ptr<TDigestCentroid>(nullptr) );
            }
            else
            {
                std::shared_ptr<TDigestCentroid> newItem(new TDigestCentroid());
                newItem->fromJson(item);
                m_Centroids.push_back( newItem );
            }
            
        }
    }
    
}


std::vector<std::shared_ptr<TDigestCentroid>>& PacketAnalyzerFlowDigestResult::getCentroids()
{
    return m_Centroids;
}

}
}
}
