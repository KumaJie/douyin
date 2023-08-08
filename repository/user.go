package repository

import "sync"

type User struct {
	ID              int64  `gorm:"primary_key;column:user_id"`
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

func (*UserDao) GetUserByName(userName string) (User, error) {
	var user User
	// 可能出现NotFind异常
	if err := db.Where("username = ?", userName).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (*UserDao) GetUserByID(userID int64) (User, error) {
	var user User
	if err := db.Where("user_id = ?", userID).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
