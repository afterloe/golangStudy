package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	r "./routers"
)

func StartServer(addr *string) {
	gin.DisableConsoleColor()
	routers := gin.Default()
	r.Execute(routers)
	server := &http.Server{
		Addr: *addr,
		Handler: routers,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
