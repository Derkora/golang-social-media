package main

import (
	"golang-social-media/config"
	"golang-social-media/routes"
	"golang-social-media/models"

	"github.com/gin-gonic/gin"
	"io"
	"os"
	"log"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Like{}, &models.Comment{})

	err := config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Like{}, &models.Comment{})
	if err != nil {
		log.Fatal("Gagal migrasi:", err)
	}

	logFile, err := os.Create("gin.log")
	if err != nil {
		panic("Could not create log file")
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	routes.SetupRoutes(r)

	r.Run(":8080")
}
