构建c++ grpc客户端
===
生成过程参考[官方文档](https://grpc.org.cn/docs/languages/cpp/quickstart/)
## 环境准备
```shell
apt install cmake

export MY_INSTALL_DIR=$HOME/.local # 选择一个目录来保存本地安装的软件包, 我用的是当前用户的家目录下的.local
mkdir -p $MY_INSTALL_DIR # 创建目录
export PATH="$MY_INSTALL_DIR/bin:$PATH" # 将bin目录放出去，让其能够调用grpc的一些包
```