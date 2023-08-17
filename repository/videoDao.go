package repository

import (
	"fmt"
	"github.com/KumaJie/douyin/util"
	"log"
	"sync"
	"time"
)

// Video 模型结构体
type Videos struct {
	VideoID    int       `gorm:"primaryKey" json:"video_id"`
	UserID     int       `json:"user_id"`
	PlayURL    string    `json:"play_url"`
	CoverURL   string    `json:"cover_url"`
	Title      string    `json:"title"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

var (
	videoDao *VideoDAO
	postOnce sync.Once
)

func (Videos) TableName() string {
	return "video"
}

type VideoDAO struct {
}

// 获取按投稿时间倒序的视频列表
func (*VideoDAO) GetVideoList() ([]Videos, error) {
	var videos []Videos
	result := util.DB.Order("create_time desc").Limit(30).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(videos)
	return videos, nil
}

func (*VideoDAO) InsertVideo(video Videos) error {
	result := util.DB.Create(video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 查询用户的投稿列表
func (*VideoDAO) GetVideoListById(userId string) ([]Videos, error) {
	var videos []Videos
	err := db.Find(&videos, "user_id = ?", userId).Error
	if err != nil {
		// 记录错误日志
		log.Printf("查询视频列表失败: %v", err)
		// 返回错误信息
		return videos, err

	}
	return videos, nil
}
