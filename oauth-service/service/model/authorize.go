package model

import "github.com/jinzhu/gorm"

type Authorize struct {
	ID          int    `gorm:"primary_key" json:"id"`
	AppId       string `json:"app_id"`
	RedirectUri string `json:"redirect_uri"`
	CreateTime  string `json:"create_time"`
}

func GetAuthorize(AppId string) (Authorize, error) {

	var authorize Authorize
	err := db.Select("id, app_id, redirect_uri, create_time").Where(Authorize{AppId: AppId}).First(&authorize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return authorize, err
	}

	return authorize, nil

}
