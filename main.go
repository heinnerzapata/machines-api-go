package main

import (
	"fmt"
	"log"
	"os"

	"machines-api-go/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	fmt.Println("Checking env files...")
	err := godotenv.Load("local.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	fmt.Println("Env files OK!")
}

func loadDatabase() {
	fmt.Println("Loading Database...")
}

func serveApplication() {
	router := gin.Default()

	envPort := fmt.Sprintf(":%v", os.Getenv("PORT"))

	// Auth paths
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	router.Run(envPort)
	fmt.Println(fmt.Sprintf("Server running on port %v", envPort))
}
