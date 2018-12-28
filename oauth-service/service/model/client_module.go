package model

import "github.com/jinzhu/gorm"

type ClientModule struct {
	ID         int    `gorm:"primary_key" json:"id"`
	ModuleName string `json:"module_name"`
	Times      int    `json:"times"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

func CheckModule(ModuleName string) (bool, error) {

	var oauth ClientModule
	err := db.Select("id").Where(ClientModule{ModuleName: ModuleName, Status: 1}).First(&oauth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if oauth.ID > 0 {
		return true, nil
	}

	return false, nil

}
