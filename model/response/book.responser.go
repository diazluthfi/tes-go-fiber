package response

import (
	"time"

	"gorm.io/gorm"
)

type BookRespone struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Cover     string         `json:"cover"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
