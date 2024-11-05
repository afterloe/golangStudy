Gin封装
===
代码及调用过程可参考[main.go](main.go),[enhance_context](internal/enhance_context.go), [internal/gateway](internal/gateway)
## 前言

以往对Gin的加强是在创建一个中间键进行，代码如下
```
*gin.Engine.Use(handle HandlerFunc)
```
当多个中间件存在的时候会产生一系列无法预估的bug，这里提供一种基于Gin的封装，将对应的方法加在Context上，以减少对中间件的依赖，同时亦能保留Gin的原始能力

## 实践过程
### 编写自定义Context
```golang
// Context 封装的Context
type Context struct {
	*gin.Context
}

// HandlerFunc 路由接收的默认参数
type HandlerFunc func(c *Context)

// EnhanceContext 增强Context防范
func EnhanceContext(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&Context{c})
	}
}
```
定义路由处理方法
```go
func Login(ctx *internal.Context) {
    // do something
}
```

### 编写增强方法
```go
// Success 请求成功
func (that *Context) Success(data ...interface{}) {
	var r interface{}
	if len(data) > 0 {
		r = data[0]
	}
	that.IndentedJSON(http.StatusOK, &CommonResponse[interface{}]{
		BizCode: 0,
		Result:  r,
	})
	that.Abort() // 中断链
}

// FailWithStr 请求失败 string
func (that *Context) FailWithStr(str string) {
	that.IndentedJSON(http.StatusOK, &CommonResponse[interface{}]{
		Error:   str,
		BizCode: http.StatusInternalServerError,
	})
	that.Abort() // 中断链
}
```

### 封装
```golang
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onenet/internal/util"
)

type (
	RouterGroup struct {
		gin.RouterGroup
	}

	Engine struct {
		Config     util.Config
		Background *http.Server
		*gin.Engine
	}
)
```
gin的核心能力为engine和RouterGroup，这两类会对Context有一定依赖，故对此进行封装。    
对以下方法进行重写
```go
func (that RouterGroup) POST(relativePath string, handlers ...internal.HandlerFunc) RouterGroup {
	that.RouterGroup.POST(relativePath, internal.EnhanceContext(handlers[0]))
	return that
}

func (that RouterGroup) Group(relativePath string, handlers ...internal.HandlerFunc) RouterGroup {
	if len(handlers) != 0 {
		that.RouterGroup.Group(relativePath, internal.EnhanceContext(handlers[0]))
	} else {
		that.RouterGroup.Group(relativePath)
	}
	return that
}

func (that RouterGroup) Use(handlers ...gin.HandlerFunc) RouterGroup {
	that.RouterGroup.Use(handlers...)
	return that
}
```
这三个方法在路由绑定的时候用的频率最高，重写这三个方法让其返回我们自己定义的Context

### 调用
```go
func (that *Engine) _bind() {
	that.Use(gin.Logger())
	that.Use(gin.Recovery())

	routes := &RouterGroup{
		that.RouterGroup,
	}
	routes.POST("/authorize/register", authorize.Register)
	routes.POST("/authorize/login", authorize.Login)

	authorizeRoute := routes.Group("/").Use(AuthorizeValid())
	authorizeRoute.POST("/album/list", album.List)
	authorizeRoute.POST("/album/get_by_id", album.GetById)
	authorizeRoute.POST("/album/create_one", album.CreateOne)
	authorizeRoute.POST("/album/delete_by_id", album.DeleteById)
}
```
中间件仍可以使用 Gin提供的，同时也能在路由测使用增强的Context。