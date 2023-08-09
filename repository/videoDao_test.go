package repository

import (
	"github.com/KumaJie/douyin/util"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	util.InitConfig()
	util.InitMysql()
	videoService := &VideoDAO{} // 创建 VideoService 实例
	v := Video{
		VideoID:    -1,
		UserID:     -1,
		PlayURL:    "...",
		CoverURL:   "..",
		Title:      ".",
		CreateTime: time.Now()}
	videoService.InsertVideo(v)

}
