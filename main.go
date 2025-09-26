package main

import (
	"bioskop-app/config"
	"bioskop-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi database
	config.ConnectDatabase()

	r := gin.Default()

	// routes
	r.POST("/bioskop", handlers.CreateBioskop)
	r.GET("/bioskop", handlers.GetBioskop)
	r.GET("/bioskop/:id", handlers.GetBioskopByID)
	r.PUT("/bioskop/:id", handlers.UpdateBioskop)
	r.DELETE("/bioskop/:id", handlers.DeleteBioskop)

	r.Run(":8080")
}
