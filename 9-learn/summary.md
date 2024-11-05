第九天
===
golang 构建rpc服务， 过程可参考[文档](https://grpc.org.cn/docs/languages/go/)

## 构建基础环境
### 安装protoc
```shell
apt install -y protoc

protoc --version
libprotoc 25.1
```
或者从github官网下载源码进行编译

### 编写proto文件
```protobuf
syntax = "proto3";

option go_package = "./;pb";

package Logic;

service UserServer {
  rpc Login(LoginRequest) returns (LoginResponse);
}

message LoginRequest {
  string LoginName = 1;
  string Scrip = 2;
}

message LoginResponse {
  string SessionID = 1;
  string AccessToken = 2;
  string RefreshToken =3;
}
```
* syntax: 制定protobuf版本  
* option go_package 表示生成的代码存放位置
> 具体编写内容可参考[这里](https://developers.google.com/protocol-buffers)

### 安装go的proto插件
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# 添加到path中，便于proto查询插件
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 使用插件生成pb文件
```shell
protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto
```
* proto_path proto文件存放的目录
* go_out     输出文件存放位置 

```shell
ls pb
user_service_grpc.pb.go  user_service.pb.go
```
执行成功后会有两个文件

## 开发gRPC服务
### 实现proto文件中定义的服务

```go
package main

import (
	"GoGrpcDemo/pb" // proto 文件生成pb文件的目录 引用进来
	"errors"
)

type userServerImpl struct {
	pb.UnimplementedUserServerServer
}

// Login 实现proto文件中定义的方法
func (*userServerImpl) Login(...interface{}) (interface{}, error) {
	// TODO
	return nil, errors.New("not impl")
}
```
编写gRPC服务
```go
package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    "GoGrpcDemo/pb"
    "GoGrpcDemo/logic"
)

func main() {
	listen, err := net.Listen("tcp", ":9092") // 监听地址
	if err != nil {
		log.Panic(err)
	}
	s := grpc.NewServer() // 获取grpc服务实例
	pb.RegisterUserServerServer(s, &userServerImpl{})  // 注册proto服务

    // 启动服务后需要关闭
	defer func() {
		s.Stop()
		err := listen.Close()
		if err != nil {
			log.Panic(err)
			return
		}
	}()

	log.Printf("Listening on port 9092 \n")
	err = s.Serve(listen)  // 启动监听
	if err != nil {
		log.Panic(err)
	}
}
```

## 构建及测试
### 启动gRPC服务
```shell
go build -v .

./GoGrpcDemo                  
2024/11/05 11:41:36 Listening on port 9092
```

### 启动C++ gRPC客户端
```shell
./CppGrpcClient
Begin to readlocalhost:9092
Login success: 177150ca-d397-4279-a091-61c8ba16a904 	a32a785e-fbbe-4a78-ab6b-9d5589f20882
```

### gRPC日志
```shell
2024/11/05 11:41:36 Listening on port 9092 
2024/11/05 11:43:06 Received request: LoginName:"afterloe"  Scrip:"111111"
```