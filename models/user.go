package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// Relasi One user to many article
	Articles []Article
	Username string
	Fullname string
	Email    string
	Password string
	SocialId string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
}
