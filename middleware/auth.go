package middleware

import (
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			// 如果Token为空，则返回状态码401和错误消息”Missing token“
			c.JSON(http.StatusUnauthorized, models.Response{
				StatusCode: 1,
				StatusMsg:  "Token缺失",
			})
			c.Abort()
			return
		}
		_, err := util.GetToken(token)
		if err != nil {
			_, err := util.VerifyToken(token)
			if err != nil {
				// 如果 Token 无效或解析失败，返回状态码 401 和错误信息 "Invalid token"
				c.JSON(http.StatusUnauthorized, models.Response{
					StatusCode: 1,
					StatusMsg:  "Token过期",
				})
				c.Abort()
				return
			}
		} else {
			//	token延期
			if err := util.RenewToken(token); err != nil {
				c.JSON(http.StatusUnauthorized, models.Response{
					StatusCode: 1,
					StatusMsg:  "Token延期失败",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
