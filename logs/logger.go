package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 计算耗时
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 记录日志
		log.Printf("[%s] %s %s %v",
			c.Request.Method,   // HTTP方法
			c.Request.URL.Path, // 请求路径
			c.ClientIP(),       // 客户端IP
			latency)            // 请求耗时
	}
}
