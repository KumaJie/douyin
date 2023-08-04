package main

import (
	"github.com/KumaJie/douyin/middleware"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", func(c *gin.Context) {
		uidStr, _ := c.GetPostForm("userid")
		username, _ := c.GetPostForm("username")
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		token, _ := util.GenerateToken(uid, username, time.Hour)
		c.JSON(http.StatusOK, token)
	})
	r.GET("post", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.RedisInit(); err != nil {
		return err
	}
	return nil
}
