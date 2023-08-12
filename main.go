package main

import (
	"github.com/KumaJie/douyin/controller"
	"github.com/KumaJie/douyin/middleware"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/util"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	userCtrl := &controller.UserController{}
	r.POST("/douyin/user/register/", userCtrl.Register)
	r.POST("/douyin/user/login/", userCtrl.Login)
	r.GET("/douyin/user/", middleware.AuthMiddleware(), userCtrl.UserInfo)
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
