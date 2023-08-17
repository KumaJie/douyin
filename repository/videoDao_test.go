package repository

import (
	"fmt"
	"github.com/KumaJie/douyin/util"
	"testing"
)

func TestAddUser(t *testing.T) {
	util.InitConfig()
	util.InitMysql()
	videoService := &VideoDAO{} // 创建 VideoService 实例
	/*	v := Videos{
		VideoID:    -1,
		UserID:     -1,
		PlayURL:    "...",
		CoverURL:   "..",
		Title:      ".",
		CreateTime: time.Now()}*/
	videoList, err := videoService.GetVideoList()
	// 输出视频列表
	if err != nil {
		fmt.Println(err)
	}
	for _, video := range videoList {
		fmt.Println(video)
	}
	//videoService.InsertVideo(v)

}
