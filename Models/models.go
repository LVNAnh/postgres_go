package Models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:255;not null" json:"username"`
	StudentID uint      `gorm:"not null" json:"student"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type CreditCard struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint      `gorm:"not null;index;foreignKey:UserID" json:"user_id"`
	CardNumber string    `gorm:"size:255;not null;unique" json:"card_number"`
	CardHolder string    `gorm:"size:255;not null" json:"card_holder"`
	ExpiryDate time.Time `gorm:"not null" json:"expiry_date"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
