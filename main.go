package main

import (
	"fmt"
	"loco-assignment/controllers"
	"loco-assignment/db"
	"loco-assignment/services"
	"loco-assignment/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load dotenv")
	}

	client := db.PostgresClient{}
	client.InitializeClient()

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("numericString", utils.NumericString)
	}

	transactionService := services.NewTrasactionService(&client)

	router := gin.Default()
	router.Use(gin.Recovery())

	transactionController := controllers.NewTrasactionController(transactionService)
	transactionController.RegisterRoutes(router)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
