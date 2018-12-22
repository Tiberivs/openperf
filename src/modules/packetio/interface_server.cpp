#include <memory>
#include <map>

#include <zmq.h>

#include "core/icp_core.h"
#include "swagger/v1/model/Interface.h"
#include "packetio/generic_interface.h"
#include "packetio/json_transmogrify.h"
#include "packetio/interface_api.h"
#include "packetio/interface_server.h"

namespace icp {
namespace packetio {
namespace interface {
namespace api {

const std::string endpoint = "inproc://icp_packetio_interface";

using namespace swagger::v1::model;
using json = nlohmann::json;
using generic_stack = icp::packetio::stack::generic_stack;

typedef std::map<unsigned, std::shared_ptr<Interface>> interface_map;

std::string to_string(request_type type)
{
    const static std::unordered_map<request_type, std::string> request_types = {
        { request_type::LIST_INTERFACES,        "LIST_INTERFACES"        },
        { request_type::CREATE_INTERFACE,       "CREATE_INTERFACE"       },
        { request_type::GET_INTERFACE,          "GET_INTERFACE"          },
        { request_type::DELETE_INTERFACE,       "DELETE_INTERFACE"       },
        { request_type::BULK_CREATE_INTERFACES, "BULK_CREATE_INTERFACES" },
        { request_type::BULK_DELETE_INTERFACES, "BULK_DELETE_INTERFACES" },
        { request_type::NONE,                   "UNKNOWN"                }  /* Overloaded */
    };

    return (request_types.find(type) == request_types.end()
            ? request_types.at(request_type::NONE)
            : request_types.at(type));
}

std::string to_string(reply_code code)
{
    const static std::unordered_map<reply_code, std::string> reply_codes = {
        { reply_code::OK,           "OK"           },
        { reply_code::NO_INTERFACE, "NO_INTERFACE" },
        { reply_code::BAD_INPUT,    "BAD_INPUT"    },
        { reply_code::ERROR,        "ERROR"        },
        { reply_code::NONE,         "UNKNOWN"      },  /* Overloaded */
    };

    return (reply_codes.find(code) == reply_codes.end()
            ? reply_codes.at(reply_code::NONE)
            : reply_codes.at(code));
}

static void _handle_list_interfaces_request(generic_stack& stack, json &request, json& reply)
{
    /* Grab optional user supplied filters */
    auto port_id = get_optional_key<std::string>(request, "port_id");
    auto mac_addr = get_optional_key<std::string>(request, "eth_mac_address");
    auto ipv4_addr = get_optional_key<std::string>(request, "ipv4_address");

    json jints = json::array();

    for (int id : stack.interface_ids()) {
        auto intf = stack.interface(id);
        if ((!port_id && !mac_addr && !ipv4_addr)
            || (port_id && port_id == std::to_string(intf->port_id()))
            || (mac_addr && mac_addr == intf->mac_address())
            || (ipv4_addr && ipv4_addr == intf->ipv4_address())) {
            jints.emplace_back(make_swagger_interface(*intf)->toJson());
        }
    }

    reply["code"] = reply_code::OK;
    reply["data"] = jints.dump();
}

static void _handle_create_interface_request(generic_stack& stack, json& request, json& reply)
{
    try {
        auto interface_model = json::parse(request["data"].get<std::string>()).get<Interface>();
        auto result = stack.create_interface(make_config_data(interface_model));
        if (!result) {
            throw std::runtime_error(result.error());
        }
        reply["code"] = reply_code::OK;
        reply["data"] = make_swagger_interface(
            *stack.interface(result.value()))->toJson().dump();
    } catch (const json::exception& e) {
        reply["code"] = reply_code::BAD_INPUT;
        reply["error"] = json_error(e.id, e.what());
    } catch (const std::runtime_error& e) {
        reply["code"] = reply_code::BAD_INPUT;
        reply["error"] = json_error(EINVAL, e.what());
    }
}

static void _handle_get_interface_request(generic_stack& stack, json& request, json& reply)
{
    int id = request["id"].get<int>();
    auto intf = stack.interface(id);
    if (intf) {
        reply["code"] = reply_code::OK;
        reply["data"] = make_swagger_interface(*intf)->toJson().dump();
    } else {
        reply["code"] = reply_code::NO_INTERFACE;
    }
}

static void _handle_delete_interface_request(generic_stack& stack, json& request, json& reply)
{
    stack.delete_interface(request["id"].get<int>());
    reply["code"] = reply_code::OK;
}

static void _handle_bulk_create_interface_request(generic_stack& stack, json& request, json& reply)
{
    /* Check input */
    std::vector<int> success_list;
    try {
        json response = {
            {"items", json::array()}
        };
        for (auto& item : request["items"]) {
            auto interface_model = item.get<Interface>();
            auto result = stack.create_interface(make_config_data(interface_model));
            if (!result) {
                throw std::runtime_error(result.error());
            }
            success_list.push_back(result.value());
            response["items"].emplace_back(make_swagger_interface(
                                               *stack.interface(result.value()))->toJson());
        }
        reply["code"] = reply_code::OK;
        reply["data"] = response.dump();
        return;
    } catch (const json::exception& e) {
        reply["code"] = reply_code::BAD_INPUT;
        reply["error"] = json_error(e.id, e.what());
    } catch (const std::runtime_error& e) {
        reply["code"] = reply_code::BAD_INPUT;
        reply["error"] = json_error(EINVAL, e.what());
    }

    /* If we get here, then we encountered an error; clean up created interfaces */
    for (int id : success_list) {
        stack.delete_interface(id);
    }
}

static void _handle_bulk_delete_interface_request(generic_stack& stack, json& request, json& reply)
{
    for (std::string request_id : request["ids"]) {
        /* id should be a number; ignore if not */
        if (int id = strtol(request_id.c_str(), nullptr, 10);
            id != 0 || request_id == "0") {  /* strtol returns 0 on error, so make sure the user
                                              * really wants to delete 0 before doing so */
            stack.delete_interface(id);
        }
    }
    reply["code"] = reply_code::OK;
}

static int _handle_rpc_request(const icp_event_data *data, void *arg)
{
    generic_stack& stack = *(reinterpret_cast<generic_stack *>(arg));
    int recv_or_err = 0;
    int send_or_err = 0;
    zmq_msg_t request_msg;
    if (zmq_msg_init(&request_msg) == -1) {
        throw std::runtime_error(zmq_strerror(errno));
    }
    while ((recv_or_err = zmq_msg_recv(&request_msg, data->socket, ZMQ_DONTWAIT)) > 0) {
        std::vector<uint8_t> request_buffer(static_cast<uint8_t *>(zmq_msg_data(&request_msg)),
                                            static_cast<uint8_t *>(zmq_msg_data(&request_msg)) + zmq_msg_size(&request_msg));

        json request = json::from_cbor(request_buffer);
        request_type type = request.find("type") == request.end() ? request_type::NONE : request["type"].get<request_type>();
        json reply;

        switch (type) {
        case request_type::GET_INTERFACE:
        case request_type::DELETE_INTERFACE:
            ICP_LOG(ICP_LOG_TRACE, "Received %s request for interface %d\n",
                    to_string(type).c_str(),
                    request["id"].get<int>());
            break;
        default:
            ICP_LOG(ICP_LOG_TRACE, "Received %s request\n", to_string(type).c_str());
        }

        switch (type) {
        case request_type::LIST_INTERFACES:
            _handle_list_interfaces_request(stack, request, reply);
            break;
        case request_type::CREATE_INTERFACE:
            _handle_create_interface_request(stack, request, reply);
            break;
        case request_type::GET_INTERFACE:
            _handle_get_interface_request(stack, request, reply);
            break;
        case request_type::DELETE_INTERFACE:
            _handle_delete_interface_request(stack, request, reply);
            break;
        case request_type::BULK_CREATE_INTERFACES:
            _handle_bulk_create_interface_request(stack, request, reply);
            break;
        case request_type::BULK_DELETE_INTERFACES:
            _handle_bulk_delete_interface_request(stack, request, reply);
            break;
        default:
            reply["code"] = reply_code::ERROR;
            reply["error"] = json_error(ENOSYS, "Request type not implemented in packetio interface server");
        }

        std::vector<uint8_t> reply_buffer = json::to_cbor(reply);
        if ((send_or_err = zmq_send(data->socket, reply_buffer.data(), reply_buffer.size(), 0))
            != static_cast<int>(reply_buffer.size())) {
            ICP_LOG(ICP_LOG_ERROR, "Request reply failed: %s\n", zmq_strerror(errno));
        } else {
            ICP_LOG(ICP_LOG_TRACE, "Sent %s reply to %s request\n",
                    to_string(reply["code"].get<reply_code>()).c_str(),
                    to_string(type).c_str());
        }
    }

    zmq_msg_close(&request_msg);

    return (((recv_or_err < 0 || send_or_err < 0) && errno == ETERM) ? -1 : 0);
}

server::server(void* context,
               icp::core::event_loop& loop,
               generic_stack& stack)
    : m_socket(icp_socket_get_server(context, ZMQ_REP, endpoint.c_str()))
{
    struct icp_event_callbacks callbacks = {
        .on_read = _handle_rpc_request
    };
    loop.add(m_socket.get(), &callbacks, &stack);
}

}
}
}
}