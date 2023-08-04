package middleware

import (
	"github.com/KumaJie/douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		switch c.Request.Method {
		case "GET":
			token = c.Query("token")
		case "POST":
			token = c.PostForm("token")
		}
		if token == "" {
			// 如果Token为空，则返回状态码401和错误消息”Missing token“
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		_, err := util.VerifyToken(token)
		if err != nil {
			// 如果 Token 无效或解析失败，返回状态码 401 和错误信息 "Invalid token"
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
