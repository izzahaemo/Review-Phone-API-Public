package models

type (
	// Submenu
	Submenu struct {
		ID     uint   `gorm:"primary_key" json:"id"`
		MenuID uint   `json:"menuID"`
		Title  string `json:"title"`
		Url    string `json:"url"`
		Icon   string `json:"icon"`
		Menu   Menu   `json:"-"`
	}
)
