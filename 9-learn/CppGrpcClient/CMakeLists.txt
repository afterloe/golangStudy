cmake_minimum_required(VERSION 3.8)

project(CppGrpcDemo C CXX)
set(CMAKE_CXX_STANDARD 20)

include("/home/afterloe/.local/grpc/examples/cpp/cmake/common.cmake")

# Proto file
get_filename_component(us_proto "src/proto/user_service.proto" ABSOLUTE)
get_filename_component(us_proto_path "${us_proto}" PATH)

# Generated sources
set(us_proto_srcs "src/user_service.pb.cc")
set(us_proto_hdrs "src/user_service.pb.h")
set(us_grpc_srcs "src/user_service.grpc.pb.cc")
set(us_grpc_hdrs "src/user_service.grpc.pb.h")

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# us_grpc_proto
add_library(us_grpc_proto
        ${us_grpc_srcs}
        ${us_grpc_hdrs}
        ${us_proto_srcs}
        ${us_proto_hdrs})
target_link_libraries(us_grpc_proto
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF})

# Targets greeter_[async_](client|server)
foreach(_target main)
    add_executable(${_target} "src/${_target}.cpp")
    target_link_libraries(${_target}
            us_grpc_proto
            absl::flags
            absl::flags_parse
            ${_REFLECTION}
            ${_GRPC_GRPCPP}
            ${_PROTOBUF_LIBPROTOBUF})
endforeach()
