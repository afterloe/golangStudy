package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	r "../routers"
	"../integrate/logger"
)

func StartUpTCPServer(addr *string) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	routers := gin.New()
	routers.Use(gin.Recovery())
	routers.Use(logger.Logger())
	v1 := routers.Group("/v1")
	r.Execute(v1)
	server := &http.Server{
		Addr: *addr,
		Handler: routers,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
