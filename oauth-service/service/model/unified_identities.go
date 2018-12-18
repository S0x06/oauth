package model

import "github.com/jinzhu/gorm"

type UnifiedIdentities struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	AppId      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
	Desc       string `json:"desc"`
	Status     string `json:"status"`
	Type       int    `json:"type"`
	CreateTime string `json:"create_time"`
}

func CheckIdentity(appid string, appsecret string) (bool, error) {

	var identity UnifiedIdentities
	err := db.Select("id").Where(UnifiedIdentities{AppId: appid, AppSecret: appsecret}).First(&identity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if identity.ID > 0 {
		return true, nil
	}

	return false, nil

}
