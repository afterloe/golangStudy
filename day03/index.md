# 第三天

interface 和 struct 一样是能够进行循环嵌入的。嵌入的方式如下
```golan
type vo1 interface {
	equal()
}

// 嵌入式 interface 理解为 interface 的继承
type vo2 interface {
	vo1
	say(vo1) string // 也可以使用 函数传参的形式
}
```
使用的时候，和普通的使用方式是一样的，所以具有和前的扩展性[代码在这](./Main.go)
    
反射是使用`reflect`包，使用方法`ValueOf`和`TypeOf`来反射获取数据的类型和值，当然要求就是被反射的值是能够修改的。interface的值是不能进行修改的，会报错  
修改值的方法如下
```golan
  p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(3.14)
	//v.SetString("afterloe") // 类型必须是一致的才可以修改
	getType(x)
```
[代码在这](./Main1.go)

使用`runtime`可以进行go的并发调用[代码在这](./Main2.go) 使用`go`关键字，如果要运用多核的话，需要使用`GOMAXPROCS`来设置同时运行的系统线程最大数量  

多核或者多并发可以使用`channel`来实现夸线程通讯，使用关键字`chan`来执行通知的管道流，而且channel可以指定是使用普通还是缓冲模式[代码在这](./Main3.go)，当然也可以使用`select`来选择的输入channel或来设置timeout[代码在这](./Main4.go) 他的高阶使用可以再项目中进一步的开发
