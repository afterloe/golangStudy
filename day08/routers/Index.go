package routers

import "github.com/gin-gonic/gin"

func Execute(route *gin.Engine)  {
	route.GET("/", Home)
}
