package main

import (
	"dataphone/config"
	"dataphone/docs"
	"dataphone/routes"
	"dataphone/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name   Muhammad Izzah Aeman
// @contact.url    http://www.izzahaemo.my.id
// @contact.email  izzah.aemo@gmail.com

func main() {
	// for load godotenv
	// for env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Review Phone Restful Api by Muh Izzah A (made by golang)"
	docs.SwaggerInfo.Description = "This project is creating a restful API by theme 'Review Phone' Using Golang (check another project by me in www.izzahaemo.my.id)"
	docs.SwaggerInfo.Version = "1.2"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:7000")
	docs.SwaggerInfo.Schemes = []string{"http"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run(":7000")
}
