package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	r "../routers"
)

func StartUpTCPServer(addr *string) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	routers := gin.Default()
	r.Execute(routers)
	server := &http.Server{
		Addr: *addr,
		Handler: routers,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
