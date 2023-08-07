package repository

import "sync"

type User struct {
	Id              int64  `gorm:"primary_key;column:user_id"`
	Name            string `gorm:"column:username"`
	Password        string `gorm:"column:password"`
	Avatar          string `gorm:"column:avatar"`
	BackGroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func UserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (*UserDao) VerifyUser(user *User) error {
	var userInDB User
	// 可能出现NotFind异常
	if err := db.First(&userInDB, user).Error; err != nil {
		return err
	}
	user = &userInDB
	return nil
}
