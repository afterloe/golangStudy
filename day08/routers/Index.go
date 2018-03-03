package routers

import (
	"github.com/gin-gonic/gin"
	"./images"
)

func Execute(route *gin.RouterGroup) {
	route.GET("/", Home)
	route.GET("/images", images.ListImage)
	route.POST("/images", Home)
}
