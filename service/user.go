package service

import (
	"fmt"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/util"
)

type UserService struct {
}

func (*UserService) CreateUser(userName string, password string) (int64, error) {
	encrypted, _ := util.HashPwd(password)
	// 保证userName唯一
	_, err := repository.UserDaoInstance().GetUserByName(userName)
	if err == nil {
		return -1, fmt.Errorf("%v已被注册", userName)
	}

	user := &repository.User{
		Name:     userName,
		Password: string(encrypted),
	}
	if err := repository.UserDaoInstance().CreateUser(user); err != nil {
		return -1, err
	}
	return user.ID, nil
}

func (*UserService) VerifyUser(userName string, password string) (int64, error) {
	user, err := repository.UserDaoInstance().GetUserByName(userName)
	if err != nil {
		return -1, fmt.Errorf("用户%v不存在", userName)
	}
	if !util.VerifyPwd(user.Password, password) {
		return -1, fmt.Errorf("密码错误")
	}
	return user.ID, nil
}

func (*UserService) GetUserInfo(userID int64) (repository.User, error) {
	user, err := repository.UserDaoInstance().GetUserByID(userID)
	if err != nil {
		return repository.User{}, err
	}
	return user, nil
}
