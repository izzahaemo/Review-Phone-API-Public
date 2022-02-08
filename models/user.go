package models

import (
	"html"
	"strings"
	"time"

	"dataphone/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	// User
	User struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Username  string    `gorm:"not null;unique" json:"username"`
		Email     string    `gorm:"not null;unique" json:"email"`
		Password  string    `gorm:"not null;" json:"password"`
		Pict      string    `json:"Pict"`
		RoleID    uint      `json:"roleID"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Reviews   []Review  `json:"review"`
		Role      Role      `json:"-"`
	}
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (int, int, string, error) {

	var err error
	var role int = 0
	var iduser int = 0

	u := User{}

	err = db.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return iduser, role, "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return iduser, role, "", err
	}
	role = int(u.RoleID)
	iduser = int(u.ID)
	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return iduser, role, "", err
	}
	return iduser, role, token, nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
