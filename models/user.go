package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// Relasi One user to many article
	Articles []Article
	NIS      string //`gorm:"unique"`
	Username string //`gorm:"size:100; unique"`
	Fullname string `gorm:"size:100"`
	Email    string `gorm:"size:100"`
	Password string
	SocialID string
	Provider string `gorm:"size:50"`
	Avatar   string
	Role     string `gorm:"size:50"`
}
