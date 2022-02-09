package config

import (
	"dataphone/models"
	"dataphone/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")
		// production
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
		db.AutoMigrate(&models.Phone{}, &models.Brand{})
		db.AutoMigrate(&models.Review{}, &models.Phone{}, &models.User{})
		db.AutoMigrate(&models.User{}, &models.Role{Name: "Admin"})
		db.AutoMigrate(&models.Accessmenu{}, &models.Role{}, &models.Menu{})
		db.AutoMigrate(&models.Submenu{}, &models.Menu{})

		return db
	} else {
		//DEVELOPMENT
		username := "root"
		password := ""
		host := "tcp(127.0.0.1)"
		database := "data_phone"

		dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}
		db.AutoMigrate(&models.Phone{}, &models.Brand{})
		db.AutoMigrate(&models.Review{}, &models.Phone{}, &models.User{})
		db.AutoMigrate(&models.User{}, &models.Role{})
		db.AutoMigrate(&models.Accessmenu{}, &models.Role{}, &models.Menu{})
		db.AutoMigrate(&models.Submenu{}, &models.Menu{})
		return db
	}
}
