package controller

import (
	"github.com/KumaJie/douyin/service"
	"github.com/KumaJie/douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userId, err := service.VerifyUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User not exist"},
		})
		return
	}
	token, err := util.GenerateToken(userId, username, time.Hour)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserId: userId,
		Token:  token,
	})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userId, err := service.CreateUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}
	token, err := util.GenerateToken(userId, username, time.Hour)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserId: userId,
		Token:  token,
	})
}
