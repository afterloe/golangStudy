package logger

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"os"
	"io"
)

var (
	out, err io.Writer
	infoLayout, logLayout, errorLayout, timeFormat string
)

func init() {
	out = os.Stdout
	err = os.Stderr
	timeFormat = "2006/01/02 - 15:04:05"
	logLayout = "[Cynomys][log][%v] - %3d | %13v | %15s | %-7s %s\n"
	infoLayout = "[Cynomys][info][%v] - %s \n"
	errorLayout = "[Cynomys][error][%v] - %s \n"
}

func GetFormatTime() string {
	return time.Now().Format(timeFormat)
}

func Error(info string) {
	fmt.Fprintf(err, errorLayout, GetFormatTime(), info)
}

func Info(info string) {
	fmt.Fprintf(out, infoLayout, GetFormatTime(), info)
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

		fmt.Fprintf(out, logLayout, end.Format(timeFormat), statusCode,
			latency, clientIP, method, path)
	}
}
