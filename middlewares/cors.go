package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 跨域设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理请求
		//c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Header("Access-Control-Allow-Origin", "*") // 允许所有IP访问
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		//c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Origin")
		//c.Header("Access-Control-Allow-Credentials", "true")
		//c.Header("Access-Control-Max-Age", "86400")
		method := c.Request.Method
		//// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
