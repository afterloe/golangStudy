package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../routers"
	"../integrate/logger"
	"../integrate/notSupper"
)

var notFoundStr, notSupperStr string

func init() {
	notFoundStr = "route is not defined."
	notSupperStr = "method is not supper"
}

func StartUpTCPServer(addr *string) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(logger.Logger())
	engine.NoRoute(notSupper.NotFound(&notFoundStr))
	engine.NoMethod(notSupper.NotSupper(&notSupperStr))
	routers.Execute(engine.Group("/v1"))
	server := &http.Server{
		Addr: *addr,
		Handler: engine,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
