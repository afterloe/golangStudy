package main

import (
	"github.com/gin-gonic/gin"
	"./dockerE"
)

func main() {
	r := gin.Default()

	r.GET("/imageList", func(c *gin.Context) {
		images := dockerE.GetImageList()
		c.JSON(200, gin.H{
			"context": images,
			"code": 200,
			"msg": nil,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}