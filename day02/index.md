# 第二天

> 代码链接 [./Main.go](./Main.go)

defer 延迟(defer)语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会 按照逆序执行，最后该函数返回进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的 资源，不然很容易造成资源泄露等问题

```golang
func readFile() bool {
 	file.Open("file")
 	if errX {
 		file.close()
 		return false;
 	}
 	if errY {
 		file.close()
 		return false;
 	}
 	file.close()
 	return true;
 }
```
使用`defer`优化之后
```golang
func readFile() bool {
 	file.Open("file")
 	defer file.Close() // defer 会在函数退出前调用
 	if errX {
 		return false;
 	}
 	if errY {
 		return false;
 	}
 	return true;
 }
```
如果有很多调用defer，那么defer是采用后进先出模式
```golang
func mockReadFile() {
	fmt.Println("打开文件")
	defer fmt.Println("记录日志")
	defer fmt.Println("关闭文件")
	fmt.Println("做一些事情")
}

/*
打开文件
做一些事情
关闭文件
记录日志
*/
```

同脚本语言一样，golang 也是可以将函数 作为值、类型进行传递的 通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值  
```golang
type testInt func (int, int) int // 定义一个函数类型

// 函数当做值和类型 写一些通用接口的时候非常有用
func useType(a int, b int, f testInt) (value int) {
	value = f(a, b)
	return
}
```

`for` 也可以配合 range 用于读取 splice 和 map的 数据
```golang
simpleMap := map[string]int { "one":1, "two":2, "three":3 }
	for key, value:= range simpleMap {
		fmt.Printf("simpleMap key is %s \n", key)
		fmt.Printf("simapleMap Value is %d \n", value)
	}
```
