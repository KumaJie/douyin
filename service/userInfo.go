package service

import (
	"fmt"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/util"
)

func CreateUser(userName string, password string) (int64, error) {
	encrypted, _ := util.HashPwd(password)
	user := &repository.User{
		Name:     userName,
		Password: string(encrypted),
	}
	if err := repository.UserDaoInstance().CreateUser(user); err != nil {
		return -1, err
	}
	return user.ID, nil
}

func VerifyUser(userName string, password string) (int64, error) {
	user, err := repository.UserDaoInstance().GetUserByName(userName)
	if err != nil {
		return -1, fmt.Errorf("用户不存在")
	}
	if !util.VerifyPwd(user.Password, password) {
		return -1, fmt.Errorf("密码错误")
	}
	return user.ID, nil
}

func GetUserInfo(userID int64) (repository.User, error) {
	user, err := repository.UserDaoInstance().GetUserByID(userID)
	if err != nil {
		return repository.User{}, err
	}
	return user, nil
}
