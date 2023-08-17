package service

import (
	"fmt"
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/util"
	"log"
	"time"
)

type VideoService struct {
	videoDAO *repository.VideoDAO
}

func (s *VideoService) NewVideoService() *VideoService {
	// 创建 VideoDAO 对象
	videoDAO := &repository.VideoDAO{}

	// 创建并初始化 VideoService 对象
	videoService := &VideoService{
		videoDAO: videoDAO,
	}

	return videoService
}

// GetDouyinFeed 根据请求参数获取 Douyin Feed 数据
func (s *VideoService) GetDouyinFeed() (*models.DouyinFeedResponse, error) {

	var response = &models.DouyinFeedResponse{}

	videos, err := s.videoDAO.GetVideoList()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		response.StatusCode = 1
		response.StatusMsg = "fail"
		return response, err
	}

	response.StatusCode = 0
	response.StatusMsg = "success"
	response.VideoList = videos

	if len(videos) > 0 {
		response.NextTime = videos[len(videos)-1].CreateTime.Unix()
	}

	return response, nil
}

func (s *VideoService) CreateVideo(req models.CreateVideoRequest) error {

	//获取用户id信息
	cla, err := util.VerifyToken(req.Token)
	if err != nil {
		return err
	}

	err = saveVideoToFile(req.Data, req.Title)
	if err != nil {
		return err
	}

	vid := saveVideoToAli(req.Title)
	time.Sleep(3 * time.Second)
	v, err := GetPlayInfo(vid)
	if err != nil {
		fmt.Println(err)
	}
	v1 := repository.Videos{
		VideoID:    v.VideoID,
		UserID:     int(cla.UserId),
		PlayURL:    v.PlayURL,
		CoverURL:   v.CoverURL,
		Title:      v.Title,
		CreateTime: v.CreateTime,
	}

	fmt.Println(v.PlayURL)
	err = s.videoDAO.InsertVideo(v1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// 获取投稿视频
func (s *VideoService) GetPublishList(token string, userId string) (*models.DouyinPublishListResponse, error) {

	var response = &models.DouyinPublishListResponse{}

	/*	type Video struct {
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
	*/
	//获取这个用户投稿过的视频
	videos, err := s.videoDAO.GetVideoListById(userId)
	var videoList []*models.Video // 创建一个初始大小为0的整数切片

	for _, value := range videos {
		v := &models.Video{
			ID:            int64(value.VideoID),
			PlayURL:       value.PlayURL,
			CoverURL:      value.CoverURL,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    true,
			Title:         value.Title,
		}
		videoList = append(videoList, v)
	}

	if err != nil {
		log.Println(err)
		response.StatusCode = 1
		response.StatusMsg = "fail"
		return response, err
	}

	response.StatusCode = 0
	response.StatusMsg = "success"
	response.VideoList = videoList
	return response, nil
}
