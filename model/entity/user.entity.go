package entity

type User struct {
	Name    string `json:"name" gorm:"primaryKey;type:varchar(10);not null" `
	Email   string `json:"email" gorm:"type:varchar(10);not null" `
	Address string `json:"address" gorm:"type:varchar(50);not null" `
	Phone   string `json:"phone" gorm:"type:varchar(50);not null"`
}
