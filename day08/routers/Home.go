package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, Success("data is ready"))
}