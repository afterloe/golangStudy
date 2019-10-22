# A Tour of Go
> create by (afterloe)[605728727@qq.com]  
> MIT License  
> version 1.0  

## Packages
每个Go语言由(Packages)包构成， 通过 import进行导入, 单行的可以使用`import "fmt"`进行，多个包引入可以使用如下方法 
```
import (
    "fmt"
    "math"
)
```
> 不过使用分组导入语句，效果会更好


