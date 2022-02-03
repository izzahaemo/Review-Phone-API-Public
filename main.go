package main

import (
	"dataphone/config"
	"dataphone/docs"
	"dataphone/routes"
	"dataphone/utils"
	"log"

	"github.com/joho/godotenv"
)

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
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "reviewphoneapi.project-izzah.my.id")
	docs.SwaggerInfo.Schemes = []string{"https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run(":8092")
}
