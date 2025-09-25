package models

import "gorm.io/gorm"

type Bioskop struct {
    ID     uint    `json:"id" gorm:"primaryKey"`
    Nama   string  `json:"nama" binding:"required"`
    Lokasi string  `json:"lokasi" binding:"required"`
    Rating float32 `json:"rating"`
}

func MigrateDB(db *gorm.DB) {
    db.AutoMigrate(&Bioskop{})
}
