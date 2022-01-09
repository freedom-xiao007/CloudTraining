package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义日志输出格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	v1 := router.Group("/v1")
	{
		hello := v1.Group("/hello")
		{
			hello.GET("", helloWorld)
		}
	}

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func helloWorld(ctx *gin.Context) {
	for key, values := range ctx.Request.Header {
		log.Println(key, values)
		for _, value := range values {
			ctx.Writer.Header().Set(key, value)
		}
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = "default"
	}
	ctx.Writer.Header().Set("version", version)

	ctx.JSON(http.StatusOK, "Hello")
}
