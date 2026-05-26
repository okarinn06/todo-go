package main

import (
	"log"
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println("Not found file .env")
    }

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}