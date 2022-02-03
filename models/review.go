package models

import "time"

type (
	// Review
	Review struct {
		ID         uint      `gorm:"primary_key" json:"id"`
		Isi        string    `json:"isi"`
		Created_at time.Time `json:"created_at"`
		PhoneID    uint      `json:"phoneID"`
		UserID     uint      `json:"userID"`
		Phone      Phone     `json:"-"`
		User       User      `json:"-"`
	}
)
