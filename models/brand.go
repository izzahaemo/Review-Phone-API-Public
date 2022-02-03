package models

import "gorm.io/gorm"

type (
	// Brand
	Brand struct {
		ID     uint    `gorm:"primary_key" json:"id"`
		Name   string  `json:"name"`
		Logo   string  `json:"logo"`
		Phones []Phone `json:"-"`
	}
)

func UserBrand(username string, db *gorm.DB) (bool, error) {
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

func NameBrand(idbrand string, db *gorm.DB) (string, error) {
	var err error
	var namep string
	b := Brand{}
	err = db.Model(Brand{}).Where("id = ?", idbrand).Take(&b).Error
	if err != nil {
		return "", err
	}
	namep = b.Name
	return namep, err
}
