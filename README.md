# go语言学习笔记

> create by [afterloe](lm6289511@gmail.com)  
> version is 1.3  
> MIT License    

## 目录
<a href="#note">笔记</a>  
<a href="#backup">备忘录</a>  

## golang 教程 学习笔记&开发日记
golang 开发的规范请参考[官方的标准](https://github.com/golang/go/wiki/CodeReviewComments)

### <a id="note">笔记</a>
golang day08中的项目已经单独立项了，可以关注 https://github.com/afterloe/AwPaas 下的 https://github.com/afterloe/awpaas-manager 项目
> 时隔半年继续更新，用go做了一个项目，将项目中碰到的点滴分享一下。将原来的第七天的内容改为了udp，第九天增加了c++扩展, 具体目录如下

直接阅读点[这里](./SUMMARY.md)
* [第一天](./src/1-learn/summary.md)
    * 变量、常量
    * map & slice 的初始化及使用
    * make & new 以及默认值
* [第二天](day02/index.md)
    * 函数、流程、goto、循环
    * [结构体 struct](day02/struct.md)
* [第三天](day03/index.md)
    * channel、多线程
    * 反射、匿名函数、interface继承
* [第四天](day04/index.md)
    * 反射的使用
    * 读取、写入文件与json处理
    * 分包分模块使用
* [第五天](day05/index.md)
    * 数据库连接
    * 字符串模板使用，输出流字符串转换
    * interface进阶使用
* [第六天](day06/index.md)
    * http服务之默认路由
    * 自定义路由
    * 文件上传
* [第七天](day07/index.md)
    * udp服务
    * udp客户端
    * 远程命令控制
* [第八天](day08/index.md)
    * 单元测试
    * 压力、性能测试
    * mock 测试
* [第九天](day09/index.md)
    * c++拓展
    * opencv 应用
    * 图像识别探索
* [第十天](https://github.com/afterloe/awpaas-manager)
    * [构建通用go打包镜像](https://github.com/afterloe/AwPaas/tree/master/awpaas-builder)
    * [构建打包工具打包成docker镜像](https://github.com/afterloe/awpaas-route/blob/master/Makefile)
    * [实战 - 使用gin构建web基础框架](day08_framework.md)
    * 调用docker api 构建一个SOA服务管理平台
    * docker 封装go的服务实现ci
* [第十一天（实战） 前置数据网关](https://github.com/afterloe/awpaas-route)
    * SOA前置数据网关
    * 自定义黑白名单
    * 链接docker service 获取服务列表
    * 使用docker swarm 网络实现服务负载均衡
* [第十二天（实战） Docker 增量维护](https://github.com/afterloe/awpaas-repository)
    * 链接CouchDB

## <a name="backup">备忘录</a>

### 不错的开源框架
**[NSQ](https://github.com/nsqio/nsq)**  
实时分发的消息平台，用于极大规模的数据处理，处理量级10亿+。  

**[Traefik](https://github.com/containous/traefik)**  
开源的反向代理与负载均衡工具。它最大的优点是能够与常见的微服务系统直接整合，可以实现自动化动态配置。  

**[Influxdb](https://github.com/influxdata/influxdb)**  
开源的分布式时序、时间和指标数据库，使用go语言编写，无需外部依赖  

**[Grafana](https://github.com/grafana/grafana)**  
开源的，具有丰富功能的度量标准仪表板和图形编辑器，用于显示Graphite，Elasticsearch，OpenTSDB，Prometheus和InfluxDB等数据，定制化高  

**[Fabric](https://github.com/hyperledger/fabric)**  
区块链超级账本Hyperledger Fabric实现，用于联盟链开发。  

**[Caddy](https://github.com/mholt/caddy)**  
Auto HTTPS Caddy 使用 Let’s Encrypt 让你的站点全自动变成全站HTTPS，无需任何配置。当然你想使用自己的证书也行  

**[Hugo](https://github.com/gohugoio/hugo)**  
一个静态的，可伸缩的网页生成器，宣称世界上最快的建站框架。  

**[Gogs](https://github.com/gogs/gogs)**  
一款极易搭建的自助Git服务  

**[Frp](https://github.com/fatedier/frp)**  
一个可用于内网穿透的高性能的反向代理应用，支持 tcp, udp 协议，为 http 和 https 应用协议提供了额外的能力，且尝试性支持了点对点穿透  

**[Proxypool](https://github.com/henson/proxypool)**  
采集免费的代理资源为爬虫提供有效的IP代理  

**[Syncthing](https://github.com/syncthing/syncthing)**  
一个持续不断的文件同步项目，能够在两台或者多台电脑上同步文件，使用了其独有的对等自由块交换协议，速度极快  


