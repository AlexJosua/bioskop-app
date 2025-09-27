package main

import (
	"bioskop-app/config"
	"bioskop-app/handlers"
	"bioskop-app/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi database
	config.ConnectDatabase()
	models.MigrateDB(config.DB) // migrate tabel otomatis

	r := gin.Default()

	// routes
	r.POST("/bioskop", handlers.CreateBioskop)
	r.GET("/bioskop", handlers.GetBioskop)
	r.GET("/bioskop/:id", handlers.GetBioskopByID)
	r.PUT("/bioskop/:id", handlers.UpdateBioskop)
	r.DELETE("/bioskop/:id", handlers.DeleteBioskop)

	// Railway kasih port lewat env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default kalau lokal
	}

	r.Run(":" + port)
}
