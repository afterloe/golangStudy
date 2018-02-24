package notSupper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../routers"
)

func NotSupper(msg *string) func(context *gin.Context) {
	return func (c *gin.Context) {
		c.Next()
		c.JSON(http.StatusOK, routers.Fail(http.StatusMethodNotAllowed, *msg))
	}
}