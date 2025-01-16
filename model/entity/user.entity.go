package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(10);not null"`
	Email     string         `json:"email" gorm:"type:varchar(20);index:unique;not null"`
	Address   string         `json:"address" gorm:"type:varchar(50);not null"`
	Phone     string         `json:"phone" gorm:"type:varchar(50);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
