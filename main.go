package main

import (
	"bioskop-app/config"
	"bioskop-app/handlers"
	"bioskop-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi database
	config.ConnectDatabase()

	// migrate tabel otomatis (jika belum ada di database)
	models.MigrateDB(config.DB)

	r := gin.Default()

	// routes
	r.POST("/bioskop", handlers.CreateBioskop)
	r.GET("/bioskop", handlers.GetBioskop)
	r.GET("/bioskop/:id", handlers.GetBioskopByID)
	r.PUT("/bioskop/:id", handlers.UpdateBioskop)
	r.DELETE("/bioskop/:id", handlers.DeleteBioskop)

	// jalankan server di port 8080
	r.Run(":8080")
}
