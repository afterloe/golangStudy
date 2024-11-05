# go语言学习笔记

> create by [afterloe](lm6289511@gmail.com)  
> version is 1.4  
> MIT License    

## 目录
- <a href="#note">笔记</a>  
- <a href="#backup">备忘录</a>  
- <a href="#install">golang 环境准备</a>

## golang 教程 学习笔记&开发日记
golang 开发的规范请参考[官方的标准](https://github.com/golang/go/wiki/CodeReviewComments)

### <a id="note">笔记</a>
golang day08中的项目已经单独立项了，可以关注 https://github.com/afterloe/AwPaas 下的 https://github.com/afterloe/awpaas-manager 项目
> 时隔半年继续更新，用go做了一个项目，将项目中碰到的点滴分享一下。将原来的第七天的内容改为了udp，第九天增加了c++扩展, 具体目录如下

直接阅读点[这里](./SUMMARY.md)

* [第一天](1-learn/summary.md)
    * 变量、常量
    * map & slice 的初始化及使用
    * make & new 以及默认值
* [第二天](2-learn/summary.md)
    * 函数、流程、goto、循环
    * [结构体 struct](2-learn/struct.md)
* [第三天](3-learn/summary.md)
    * channel、多线程
    * 反射、匿名函数、interface继承
* [第四天](4-learn/summary.md)
    * 反射的使用
    * 读取、写入文件与json处理
    * 分包分模块使用
* [第五天](5-learn/summary.md)
    * 数据库连接
    * 字符串模板使用，输出流字符串转换
    * interface进阶使用
* [第六天](6-learn/summary.md)
    * http服务之默认路由
    * 自定义路由
    * 文件上传
* [第七天](7-learn/summary.md)
    * udp服务
    * udp客户端
    * 远程命令控制
* [第八天](8-learn/summary.md)
    * 单元测试
    * 压力、性能测试
    * mock 测试
* [第九天](9-learn/summary.md)
    * gRPC服务构建及测试
* [第十天](10-learn/summary.md)
    * [构建通用Go打包镜像](10-learn/build_by_docker.md)

## <a name="backup">备忘录</a>

## <a name="install">golang 环境准备</a>
```bash
wget https://dl.google.com/go/go1.12.linux-amd64.tar.gz
sudo tar xzvf go1.12.linux-amd64.tar.gz -C /usr/local/
```

配置go path 及 go root
```bash
$ sudo vim /etc/profile.d/go.sh
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# Set the GOPROXY environment variable
export GOPROXY=https://goproxy.io,direct
# Set environment variable allow bypassing the proxy for specified repos (optional)
export GOPRIVATE=git.mycompany.com,github.com/my/private
```
