package integrate

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"os"
	"io"
)

var out io.Writer

func init() {
	out = os.Stdout
}

func Logger() func(*gin.Context) {

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if "" != raw {
			path = path + "?" + raw
		}

		fmt.Fprintf(out, "[Cynomys] %v | %3d | %13v | %15s | %-7s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode, latency, clientIP, method, path)
	}
}
