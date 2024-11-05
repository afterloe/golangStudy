构建c++ grpc客户端
===
生成过程参考[官方文档](https://grpc.org.cn/docs/languages/cpp/quickstart/)  
源码位置: [CppGrpcClient](./CppGrpcClient/src/main.cpp)

## 环境准备
构建环境
```shell
apt install cmake

export MY_INSTALL_DIR=$HOME/.local # 选择一个目录来保存本地安装的软件包, 我用的是当前用户的家目录下的.local
mkdir -p $MY_INSTALL_DIR # 创建目录
export PATH="$MY_INSTALL_DIR/bin:$PATH" # 将bin目录放出去，让其能够调用grpc的一些包

sudo apt install -y build-essential autoconf libtool pkg-config # 其他的一些包
```

准备源码
```shell
git clone --recurse-submodules -b v1.62.0 --depth 1 --shallow-submodules https://github.com/grpc/grpc
```

源码构建
```shell
cd grpc
mkdir -p cmake/build
pushd cmake/build
cmake -DgRPC_INSTALL=ON \
    -DgRPC_BUILD_TESTS=OFF \
    -DCMAKE_INSTALL_PREFIX=$MY_INSTALL_DIR \
    ../..
make -j 4
make install
popd
```

## 编写client

### 引入编写好的proto文件
```shell
mkdir CppGrpcDemo && cd CppGrpcDemo
mkdir -p src/proto
cp ${YOU_PATH_PROTO}/*.proto src/proto
```

### 编写Makefile
```shell
clean:
	rm -rf build

proto:
	rm -rf src/user_service.*
	@protoc -I src/proto --grpc_out=src --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` src/proto/*.proto
	@protoc -I src/proto --cpp_out=src src/proto/*.proto
	echo 'generator file success, see src/*'

.PHONY: proto clean
```
执行 `make proto`

### 编写CMakeLists.txt
```
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
```

### 编写客户端cpp文件
```c++
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

```

## 编译
```shell
cd CppGrpcDemo
rm -rf build && mkdir build
cd build
cmake ..
make

[ 20%] Building CXX object CMakeFiles/us_grpc_proto.dir/src/user_service.grpc.pb.cc.o
[ 40%] Building CXX object CMakeFiles/us_grpc_proto.dir/src/user_service.pb.cc.o
[ 60%] Linking CXX static library libus_grpc_proto.a
[ 60%] Built target us_grpc_proto
[ 80%] Building CXX object CMakeFiles/main.dir/src/main.cpp.o
[100%] Linking CXX executable main
[100%] Built target main
```
直接运行 `./main` 即可