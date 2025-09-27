package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := os.Getenv("DATABASE_URL") // ambil dari Railway
    if dsn == "" {
        log.Fatal("DATABASE_URL tidak ditemukan")
    }

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal koneksi ke database:", err)
    }

    DB = database
    fmt.Println("✅ Berhasil terkoneksi ke database")
}
