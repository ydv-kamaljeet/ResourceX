package main

import (
	"fmt"
	"log"
	"os"

	"book.com/internal/db"
	"book.com/internal/routes"
	"book.com/internal/storage"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //have pre-existing logger middleware

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	db.Init()
	if db.DB == nil {
		log.Fatal("Db not initialized")
	}
	fmt.Println("App started successfully after DB connection.")

	// err := db.DB.AutoMigrate(&models.Book{})
	// if err != nil {
	// 	log.Fatal("Failed to migrate database schema")
	// }
	storage.InitAzureBlob() // ✅ Init Azure connection

	// Routes
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
