package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	// Data Source Name (DSN)
	const MYSQL = "root:123@tcp(127.0.0.1:3306)/tes_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi ke database
	DB, err = gorm.Open(mysql.Open(MYSQL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successful!")
}
