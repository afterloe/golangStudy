package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	r "../routers"
	log "../integrate"
)

func StartUpTCPServer(addr *string) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	routers := gin.New()
	routers.Use(gin.Recovery())
	routers.Use(log.Logger())
	v1 := routers.Group("/v1")
	r.Execute(v1)
	server := &http.Server{
		Addr: *addr,
		Handler: routers,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
