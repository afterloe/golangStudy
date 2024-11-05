package gateway

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onenet/internal"
	"onenet/internal/util/token"
	"strings"
)

func AuthorizeValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &internal.Context{Context: c}
		bearerToken := ctx.GetHeader("Authorization")
		if len(strings.Split(bearerToken, " ")) != 2 {
			ctx.IndentedJSON(http.StatusUnauthorized, &internal.CommonResponse[interface{}]{
				BizCode: http.StatusUnauthorized,
				Error:   "未授权",
			})
			ctx.Abort()
			return
		}
		tokenStr := strings.Split(bearerToken, " ")[1]
		payload, err := token.Inspect.VerifyToken(tokenStr)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, &internal.CommonResponse[interface{}]{
				BizCode: http.StatusUnauthorized,
				Error:   "未授权",
			})
			ctx.Abort()
			return
		}
		err = token.ValidExpired(payload.ExpiredAt)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, &internal.CommonResponse[interface{}]{
				BizCode: http.StatusUnauthorized,
				Error:   "登陆超时",
			})
			ctx.Abort()
			return
		}
		ctx.Set("authorize", payload)
		ctx.Next()
	}
}
