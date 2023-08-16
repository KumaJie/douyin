package controller

import (
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/service"
	"github.com/KumaJie/douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService *service.UserService
}

func (u *UserController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userID, err := u.UserService.VerifyUser(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	token, err := util.GenerateToken(userID, username)
	c.JSON(http.StatusOK, models.UserLoginResponse{
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
		c.JSON(http.StatusBadRequest, models.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	token, err := util.GenerateToken(userID, username)
	c.JSON(http.StatusOK, models.UserLoginResponse{
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
		c.JSON(http.StatusBadRequest, models.Response{StatusCode: 1, StatusMsg: "Get UserInfo fail"})
		return
	}
	// 还需要增加关注数等查询
	c.JSON(http.StatusOK, models.UserInfoResponse{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		User: &models.UserInfo{
			ID:              user.ID,
			Name:            user.Name,
			FollowCount:     0,
			FollowerCount:   0,
			IsFollow:        false,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackGroundImage,
			Signature:       user.Signature,
			TotalFavorited:  0,
			WorkCount:       0,
			FavoriteCount:   0,
		},
	})

}
