package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"codechef/backend/config"
	"codechef/backend/routes"
)

func main() {
	_ = godotenv.Load()

	if err := config.ConnectMongo(); err != nil {
		log.Fatalf("mongo connect error: %v", err)
	}

	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
