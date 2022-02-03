package models

type (
	// Accessmenu
	Accessmenu struct {
		ID     uint `gorm:"primary_key" json:"id"`
		RoleID uint `json:"roleID"`
		MenuID uint `json:"menuID"`
		Role   Role `json:"-"`
		Menu   Menu `json:"-"`
	}
)
