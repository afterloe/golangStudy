package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	bizError "onenet/internal/errors"
	"onenet/internal/util/token"
	"reflect"
)

// Context 封装的Context
type Context struct {
	*gin.Context
}

// HandlerFunc 路由接收的默认参数
type HandlerFunc func(c *Context)

// EnhanceContext 增强Context方法
func EnhanceContext(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&Context{c})
	}
}

// BindAuthorizeInfo 为entity._base 中 绑定创建人、修改人id
func (that *Context) BindAuthorizeInfo(entity interface{}) {
	info := that.AuthorizeInfo()
	reflectValue := reflect.ValueOf(entity)
	reflectValue.Elem().FieldByName("CreatedBy").SetString(info.ID)
	reflectValue.Elem().FieldByName("ModifiedBy").SetString(info.ID)
}

// AuthorizeInfo 获取登陆人信息
func (that *Context) AuthorizeInfo() *token.Payload {
	val, ok := that.Get("authorize")
	if ok {
		return val.(*token.Payload)
	}
	return nil
}

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
	that.Abort()
}

// FailWithStr 请求失败 string
func (that *Context) FailWithStr(str string) {
	that.IndentedJSON(http.StatusOK, &CommonResponse[interface{}]{
		Error:   str,
		BizCode: http.StatusInternalServerError,
	})
	that.Abort()
}

// Fail 请求失败 error
func (that *Context) Fail(err error) {
	response := &CommonResponse[interface{}]{
		Error: err.Error(),
	}
	var h *bizError.BizError
	if errors.As(err, &h) {
		response.BizCode = h.ErrorCode
	} else {
		response.BizCode = http.StatusInternalServerError
	}
	that.IndentedJSON(http.StatusOK, response)
	that.Abort()
}
