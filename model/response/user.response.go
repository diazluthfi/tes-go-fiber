package response

import "time"

type UserResponse struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" `
	Email     string    `json:"email"`
	Password  string    `json:"-" gorm:"column:password"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
