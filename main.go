package main

import (
	"go-blog/config"
	"go-blog/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r := gin.Default()
	routes.Serve(r)
	r.Use(cors.New(corsConfig))
	r.Run(":" + os.Getenv("PORT"))
}
