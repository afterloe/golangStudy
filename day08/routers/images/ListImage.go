package images

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../routers"
	"../../service/docker-cli"
)

func ListImage(c *gin.Context) {
	images, err := docker_cli.ListImage()
	if nil != err {
		c.Error(err)
	}
	c.JSON(http.StatusOK, routers.Success(images))
}