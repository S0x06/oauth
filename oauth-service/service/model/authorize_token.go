package model

import "github.com/jinzhu/gorm"

type AuthorizeToken struct {
	ID           int    `gorm:"primary_key" json:"id"`
	AppId        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Code         string `json:"code"`
	GrandType    string `json:"grand_type"`
	Expiration   int    `json:"expiration"`
	RefreshTime  string `json:"refresh_time"`
	CreateTime   string `json:"create_time"`
}

func GetCode(Code string) (bool, error) {

	var authorize AuthorizeToken
	err := db.Select("id").Where(AuthorizeToken{Code: Code}).First(&authorize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if authorize.ID > 0 {
		return true, nil
	}

	return false, nil

}
