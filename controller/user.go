package controller

import (
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/service"
	"github.com/KumaJie/douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	models.Response
	UserID int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type UserInfoResponse struct {
	models.Response
	User User `json:"user"`
}

type UserController struct {
	UserService *service.UserService
}

func (u *UserController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userID, err := u.UserService.VerifyUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User not exist"},
		})
		return
	}
	token, err := util.GenerateToken(userID, username)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserID: userID,
		Token:  token,
	})
}

func (u *UserController) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userID, err := u.UserService.CreateUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User already exist"})
		return
	}
	token, err := util.GenerateToken(userID, username)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserID: userID,
		Token:  token,
	})
}

func (u *UserController) UserInfo(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	user, err := u.UserService.GetUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "Get UserInfo fail"})
		return
	}
	// 还需要增加关注数等查询
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		User: User{
			ID:              user.ID,
			Name:            user.Name,
			FollowCount:     0,
			FollowerCount:   0,
			IsFollow:        false,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackGroundImage,
			Signature:       user.Signature,
			TotalFavorited:  "",
			WorkCount:       0,
			FavoriteCount:   0,
		},
	})

}
