package service

import "github.com/KumaJie/douyin/repository"

func CreateUser(username string, password string) (userId int64, err error) {
	user := &repository.User{
		Name:     username,
		Password: password,
	}
	if err := repository.UserDaoInstance().CreateUser(user); err != nil {
		return -1, err
	}
	return user.Id, nil
}

func VerifyUser(username string, password string) (userId int64, err error) {
	user := &repository.User{
		Name:     username,
		Password: password,
	}
	if err := repository.UserDaoInstance().VerifyUser(user); err != nil {
		return -1, err
	}
	return user.Id, nil
}
