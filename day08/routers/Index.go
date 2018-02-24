package routers

import "github.com/gin-gonic/gin"

func Execute(route *gin.RouterGroup) {
	route.GET("/", Home)
}
