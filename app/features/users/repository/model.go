package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"primaryKey;type:varchar(100);unique"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
	Image    string `json:"image" gorm:"type:varchar(100);not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Image == "" {
		u.Image = "https://storage.googleapis.com/grouproject/book/images.jpeg"
	}
	return nil
}
