package models

type (
	// Role
	Role struct {
		ID          uint         `gorm:"primary_key" json:"id"`
		Name        string       `json:"name"`
		Accessmenus []Accessmenu `json:"-"`
		Users       []User       `json:"-"`
	}
)
