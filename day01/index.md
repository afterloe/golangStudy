# 第一天
> 代码链接 [./Main.go](./Main.go)

## GO 中的潜规则

Go语言有一些默认的行为。
- 大写字母开头的变量是可导出的，即其他包可以读取，是公用变量;小写字母开头的不可导出，是私有变量。
- 大写字母开头的函数也是一样，相当于class中带public关键词的公有函数;小写字母开头就是有private关键词的私有函数。

## 变量

变量分为 `var` 和 `const`。 申明变量可以有已下几种方式
```golang
var age = 2
c := 2

var (
  a = 1,
  b = 2
)
```

`:=` 是一种简写，它和`var`不能一同使用， 第二一个是可以进行批量赋值。

## map 和 slice
类似于 `HashMap` 和 `LinkedArrayList` 都有自己的默认方法。

## 各类型的初始变量

- make用于内建类型(map、slice和channel)的内存分配。new用于各种类型的内存分配，直接返回指针`*T`
- make只能创建slice、map和channel，并且返回一个有初始值(非零)的T 类型 make返回初始化后的(非零)值。

类型 | 值
---- | ----
int | 0
int8 | 0
int32 | 0
int64 | 0
uint | 0x0
rune | 0
byte | 0x0
float32 | 0
float64 | 0
boolean | false
string | ""