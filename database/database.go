package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func DatabaseInit() {
	db, err := gorm.Open("mysql", "root:123@/tes_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer db.Close() // Pastikan koneksi ditutup setelah selesai digunakan

	fmt.Println("Database connection successful!")

}
