#include <iostream>
#include <grpc++/create_channel.h>
#include <grpc++/security/credentials.h>

#include "user_service.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReader;
using grpc::ClientReaderWriter;
using grpc::ClientWriter;
using grpc::Status;
using Logic::UserServer;
using Logic::LoginRequest;
using Logic::LoginResponse;

int main(int argc, char **argv) {
    const std::string server_host = "localhost:9092";
    std::cout << "Begin to read" << server_host << std::endl;

    const std::unique_ptr<UserServer::Stub> impl = UserServer::NewStub(
        CreateChannel(server_host, grpc::InsecureChannelCredentials()));

    auto *req = new LoginRequest();
    req->set_loginname("afterloe");
    req->set_scrip("111111");

    LoginResponse response;
    impl->Login(new ClientContext(), *req, &response);
    std::cout << "Login success: " << response.accesstoken() << " \t" << response.sessionid() << std::endl;
    return 0;
}
