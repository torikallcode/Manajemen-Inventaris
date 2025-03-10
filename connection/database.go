package connection

import (
	"fmt"
	"inventaris/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	dsn := "root:torikal@tcp(127.0.0.1:3306)/inventaris?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database", err)
	}

	DB = db
	fmt.Println("Berhasil terhubung ke database")

	db.AutoMigrate(&models.User{})
	fmt.Println("Migrasi Berhasil")
}
