# 第四天

分包 通过 `import` 关键字来引入本地包[代码](./Main.go)

```golang
import (
    "./fs" // 引入本地fs包
)
```
[fs源码](./fs/Fs.go)

切记一点，小写开头的是private的，大写开头的是public的。使用`strconv`这个工具包可以经常转换比如数字转字符串、boolean转字符串。而且需要注意的是，在同一个包内是可以不用引入直接调用的，并且不能出现同名的函数。尤其是大写开头的   

在go里面，函数返回一般会带一个error
```golang
func t() (string, error) {}
```

自定义error是支持的。不过需要重写`Error()`方法。[代码在这](./fs/Error.go)

```golang
// 自定义 Error
type Error struct {
	msg string
	code int
}

// 自定义异常要实现这个方法
func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}
```

使用`io/ioutil`可以读取文件[代码](./fs/ReadFile.go) 写文件也可以这个工具[代码](./fs/WriteFile.go)    
转换JSON可以使用`encoding/json`工具包，具体使用可以[参考这里](./fs/FormatJSON.go)

反射在这里使用的很高详细的[这里](./Main.go)，使用`reflect`包就可以，`interface{}`转`map`的方法如下
```golang
	v := reflect.ValueOf(depMap) // 使用反射 获取 map

	// 获取map中的指定值
	// fmt.Println(depMap["koa"])
	fmt.Println(reflect.TypeOf(depMap)) // map[string]interface {}
	fmt.Println(v.MapIndex(reflect.ValueOf("koa"))) // 使用这个方式来直接取map 中 koa的值

	// index 第一个参数是 下标 -- 循环输出map中的值
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		fmt.Printf("%s -> %s \n", key.Interface(), val.Interface())
	}

```

`interface{}`转`slice`方法如下

```golang
	// 循环输出 slice 中的值
	keywords := obj["keywords"].([]interface{})
	fmt.Println(reflect.TypeOf(keywords))
	v = reflect.ValueOf(keywords)
	fmt.Println(v.Index(0)) // 直接取第一个数的值
	for _, val := range keywords{
		val := reflect.ValueOf(val)
		fmt.Println(val.Interface())
	}
```
