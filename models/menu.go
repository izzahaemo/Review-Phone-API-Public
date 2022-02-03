package models

type (
	// Menu
	Menu struct {
		ID          uint         `gorm:"primary_key" json:"id"`
		Name        string       `json:"name"`
		Accessmenus []Accessmenu `json:"-"`
		Submenus    []Submenu    `json:"-"`
	}
)
