package models

import (
	"gin-ranking/dao"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"addTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(username string, password string) (int, error) {
	user := User{
		Username:   username,
		Password:   password,
		AddTime:    time.Now().Unix(),
		UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func GetUsers(aid int) ([]User, error) {
	var users []User
	err := dao.Db.Where("aid = ?", aid).Order("add_time desc, id desc").First(&users).Error
	return users, err
}

func AddUserTest() (int, error) {
	user := User{Username: "test1", AddTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func SaveUser(user User) (int, error) {
	err := dao.Db.Save(&user).Error
	return user.Id, err
}

func DeleteUser(id int) error {
	err := dao.Db.Delete(&User{}, id).Error
	return err
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}
