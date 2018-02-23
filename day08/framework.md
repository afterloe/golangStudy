# 实战 - 使用gin构建web基础框架

go的架构可以说是第一次，有问题尽管开骂。我就参照以前做node的时候开发的项目进行一个大致的架构。node相关的代码也在github上

* https://github.com/afterloe/cynomy_portal_server
* https://github.com/afterloe/iot-demo
* https://github.com/afterloe/yunqi_cloud

物理架构图
```sbtshell
---- Root
 |
 | -- routers					路由
 | -- services					存放服务相关的函数
 | -- dao					存放数据相关操作的函数
 | -- config					配置文件
 | -- test					测试文件
 | -- Dockerfile				docker ci 相关的文件
 | -- Makefile					构建、ci的命令入口文件
 | -- integrate					接入外部服务
 | -- exceptions				自定义异常
 | -- start-cli.go				启动tcp服务
 | -- Main.go					主入口
 | -- package.json				ci相关配置参数
 |
```

这是基本的物理架构，可能会有修改，先放一个版本在这里
