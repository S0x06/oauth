package model

import (
	"crypto/md5"
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func GetUser(UserName string, Password string) (User, error) {

	var user User

	//	password :=
	hasPassword := md5.Sum([]byte(Password))
	md5Password := fmt.Sprintf("%x", hasPassword) //将[]byte转成16进制

	err := db.Select("id").Where(User{UserName: UserName, Password: md5Password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil

}
