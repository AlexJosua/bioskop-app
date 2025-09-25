package handlers

import (
	"bioskop-app/config"
	"bioskop-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBioskop
func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	if err := config.DB.Create(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bioskop)
}

// GetBioskop baru
func GetBioskop(c *gin.Context) {
	var bioskops []models.Bioskop

	if err := config.DB.Find(&bioskops).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}
