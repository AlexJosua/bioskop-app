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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bioskop)
}

// Get all bioskop
func GetBioskop(c *gin.Context) {
	var bioskops []models.Bioskop

	if err := config.DB.Find(&bioskops).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}

// Get bioskop by ID
func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop dengan ID tersebut tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}

// Update bioskop by ID
func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	// cek apakah data ada
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop dengan ID tersebut tidak ditemukan"})
		return
	}

	// bind data baru
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid: " + err.Error()})
		return
	}

	// validasi minimal nama & lokasi
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	// update data
	bioskop.Nama = input.Nama
	bioskop.Lokasi = input.Lokasi
	bioskop.Rating = input.Rating

	if err := config.DB.Save(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}

// Delete bioskop by ID
func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	// cek data
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop dengan ID tersebut tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
