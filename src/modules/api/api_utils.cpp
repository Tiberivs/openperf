
#include <string>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <iostream>
#include <unistd.h>
#include <string.h>

#include "api_service.h"
#include "api_internal_client.h"
#include "api/api_utils.h"

using namespace std;

namespace icp::api::utils {

static const string api_server_host          = "localhost";
static const string api_check_resource       = "/version";
static constexpr unsigned int poll_interval  = 10;  // in ms
static constexpr unsigned int max_poll_count = 6;

// Is the API port open?
static tl::expected<void, std::string> check_api_port()
{
    struct addrinfo hints = {.ai_family = AF_UNSPEC, .ai_socktype = SOCK_STREAM}, *result;

    int res = getaddrinfo(api_server_host.c_str(),
                          to_string(icp::api::api_get_service_port()).c_str(), &hints, &result);
    if (res != 0) {
        std::cerr << "Error starting up internal API client: " << gai_strerror(res) << std::endl;
        return (tl::make_unexpected("Error starting up internal API client: "
                                    + std::string(gai_strerror(res))));
    }

    unsigned int poll_count = 0;
    bool done               = false;
    for (; (poll_count < max_poll_count) && !done; poll_count++) {
        // Experiments showed that a delay is always required.
        // Sleep first, check later.
        std::this_thread::sleep_for(std::chrono::milliseconds(poll_interval * (1 << poll_count)));

        int sockfd = socket(result->ai_family, result->ai_socktype, result->ai_protocol);
        if (sockfd == -1) {
            freeaddrinfo(result);

            return (tl::make_unexpected("Error starting up internal API client: "
                                        + std::string(strerror(errno))));
        }

        if (connect(sockfd, result->ai_addr, result->ai_addrlen) == 0) { done = true; }

        close(sockfd);
    }

    freeaddrinfo(result);

    if (!done) {
        return (tl::make_unexpected(
          "Error starting up internal API client. Could not connect to API server."));
    }

    return {};
}

// Is the API able to retrieve a resource?
static tl::expected<void, std::string> check_api_resource()
{
    unsigned int poll_count = 0;
    bool done               = false;

    for (; (poll_count < max_poll_count) && !done; poll_count++) {
        auto [code, body] = icp::api::client::internal_api_get(api_check_resource);
        if (code == Pistache::Http::Code::Ok) {
            done = true;
            break;
        }

        std::this_thread::sleep_for(std::chrono::milliseconds(poll_interval * (1 << poll_count)));
    }

    if (!done) {
        return (
          tl::make_unexpected("Error starting up internal API client. Could not retrieve resource: "
                              + api_check_resource));
    }

    return {};
}

tl::expected<void, std::string> check_api_module_running()
{
    auto result = check_api_port();
    if (!result) return (result);

    result = check_api_resource();
    if (!result) return (result);

    return {};
}

}  // namespace icp::api::utils