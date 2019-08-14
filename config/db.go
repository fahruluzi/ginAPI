package config

import (
	"github.com/Fahrul-uzi/gin-full/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "zii:zii@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect Database")
	}

	DB.AutoMigrate(&models.User{})
	// menambahkan foreign key pada article, (menambahkan id user pada tabel article)
	DB.AutoMigrate(&models.Article{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// Menginisialisasi relasi
	DB.Model(&models.User{}).Related(&models.Article{})
}
