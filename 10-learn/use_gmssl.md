在go中使用国密算法
===

在国产化中数据库推荐使用达梦、人大金仓、神通等，加密算法自然是国密，Go在使用国密算法的时候需要使用如下库  
https://github.com/GmSSL/GmSSL-Go

而GmSSL是由北京大学自主开发的国产商用密码开源库，实现了对国密算法、标准和安全通信协议的全面功能覆盖，支持包括移动端在内的主流操作系统和处理器，
支持密码钥匙、密码卡等典型国产密码硬件，提供功能丰富的命令行工具及多种编译语言编程接口。
https://github.com/guanzhi/GmSSL

## 安装前准备
下载源码并进行编译和安装
```shell
cd ~/.local/src
git clone git@github.com:guanzhi/GmSSL.git

cd GmSSL
mkdir build
cd build
cmake ..
make
make test
sudo make install

cd ~/.local/lib
ls | grep gmssl
libgmssl.so  libgmssl.so.3  libgmssl.so.3.1

cd ~/.local/include
ls | grep gmssl
gmssl
```
显示如上内容表示密码库安装完毕

## Go调用
```shell
go get github.com/GmSSL/GmSSL-Go
```
若未安装国密库，该模块是无法进行安装的
```go
// _str2Ciphertext
// 国密sm3 hash
func _str2Ciphertext(str string) string {
	sm3 := gmssl.NewSm3()
	sm3.Update([]byte(str))
	cipher := sm3.Digest()
	return fmt.Sprintf("%x", cipher)
}
```
使用的是国密sm3 hash算法，可参考[internal/logic/authorize/sub_imple.go](internal/logic/authorize/sub_imple.go)的27行   

更多算法可参照GmSSL的 github。
