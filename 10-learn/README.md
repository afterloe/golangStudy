项目1 - 简单的相册管理
===
> create by [afterloe](liumin@hengzhiinfo.cn)  
> version is 1.0.15  
> MIT LICENSE

## 项目目录结构设计
| 目录/文件名称    | 说明   | 描述                                                 |
|------------|------|----------------------------------------------------|
| api        | 对外接口 | 对外提供服务的输入/输出数据结构定义。考虑到版本管理需要，往往以api/xxx/v1...存在。   |
| hack       | 工具脚本 | 存放项目开发工具、脚本等内容。例如，CLI工具的配置，各种shell/bat脚本等文件。       |
| internal   | 内部逻辑 | 业务逻辑存放目录。通过Golang internal特性对外部隐藏可见性。              |
| - cmd      | 入口指令 | 命令行管理目录。可以管理维护多个命令行。                               |
| - consts   | 常量定义 | 项目所有常量定义。                                          |
| - gateway  | 接口处理 | 接收/解析用户输入参数的入口/接口层。                                |
| - dao      | 数据访问 | 数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法         |
| - logic    | 业务封装 | 业务逻辑封装管理，特定的业务逻辑实现和封装。往往是项目中最复杂的部分。                |
| - entity   | 结构模型 | 数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义。                   |
| - service  | 业务接口 | 用于业务模块解耦的接口定义层。具体的接口实现在logic中进行注入。                 |
| - util     | 通用工具 | 用于内部调用的工具集合                                        |
| manifest   | 交付清单 | 包含程序编译、部署、运行、配置的文件。常见内容如下：                         |
| - config   | 配置管理 | 配置文件存放目录。                                          |
| - docker   | 镜像文件 | Docker镜像相关依赖文件，脚本文件等等。                             |
| - deploy   | 部署文件 | 部署相关的文件。默认提供了Kubernetes集群化部署的Yaml模板，通过kustomize管理。 |
| - protobuf | 协议文件 | GRPC协议时使用的protobuf协议定义文件，协议文件编译后生成go文件到api目录。      |
| resource   | 静态资源 | 静态资源文件。这些文件往往可以通过 资源打包/镜像编译 的形式注入到发布文件中。           |
| go.mod     | 依赖管理 | 使用Go Module包管理的依赖描述文件。                             |
| main.go    | 入口文件 | 程序入口文件。                                            |

## 数据库
CouchDB
```shell
docker run -d \
-p 5984:5984 \
--name my-couchdb \
-e COUCHDB_USER=admin \
-e COUCHDB_PASSWORD=111111hZ! \
apache/couchdb:latest
```

## 相关资料

* [增强Context, gin封装web服务](enhance_context.md)
* [国密算法集成](use_gmssl.md)
* [jwt Token](bear_token.md)
* [Swagger 配置](use_swagger.md)
* [构建通用Go打包镜像](build_by_docker.md)