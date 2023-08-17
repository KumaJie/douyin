package models

import "github.com/KumaJie/douyin/repository"

// 视频投稿
type CreateVideoRequest struct {
	Token string `json:"token" binding:"required"`
	Data  []byte `json:"data" binding:"required"`
	Title string `json:"title" binding:"required"`
}

// DouyinFeedRequest 是 Douyin Feed 请求结构体
type DouyinFeedRequest struct {
	LatestTime int64  `json:"latest_time"`
	Token      string `json:"token"`
}

// DouyinFeedResponse 是 Douyin Feed 响应结构体
type DouyinFeedResponse struct {
	StatusCode int32               `json:"status_code"`
	StatusMsg  string              `json:"status_msg"`
	VideoList  []repository.Videos `json:"video_list"`
	NextTime   int64               `json:"next_time"`
}

type Video struct {
	ID            int64 `gorm:"primaryKey"`
	Author        User  `gorm:"foreignKey:AuthorID"`
	AuthorID      int64
	PlayURL       string
	CoverURL      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
}

type User struct {
	ID              int64 `gorm:"primaryKey"`
	Name            string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
}

/*
// 获取投稿列表

	type DouyinPublishListRequest struct {
		UserId int64  `json:"user_id"`
		Token  string `json:"token"`
	}
*/
type DouyinPublishListResponse struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg,omitempty"`
	VideoList  []*Video `json:"video_list"`
}
