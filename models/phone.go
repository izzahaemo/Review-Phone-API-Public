package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	// Phone
	Phone struct {
		ID          uint      `json:"id" gorm:"primary_key"`
		Name        string    `json:"name"`
		Display     string    `json:"display"`
		Camerafront string    `json:"Camerafront"`
		Cameraback  string    `json:"Cameraback"`
		Chipset     string    `json:"chipset"`
		Ram         string    `json:"ram"`
		Battery     string    `json:"battery"`
		Pict        string    `json:"pict"`
		Created_at  time.Time `json:"created_at"`
		BrandID     uint      `json:"brandID"`
		Brand       Brand     `json:"-"`
		Reviews     []Review  `json:"-"`
	}
)

func UserPhone(username string, db *gorm.DB) (bool, error) {
	var err error
	var tru bool = true
	u := User{}
	err = db.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return tru, err
	}
	if u.RoleID == 2 {
		tru = false
		return tru, err
	}
	return tru, err
}

func NamePhone(idbrand string, db *gorm.DB) (string, error) {
	var err error
	var namep string
	p := Phone{}
	err = db.Model(Phone{}).Where("id = ?", idbrand).Take(&p).Error
	if err != nil {
		return "", err
	}
	namep = p.Name
	return namep, err
}
