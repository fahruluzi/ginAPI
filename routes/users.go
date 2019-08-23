package routes

import (
	"github.com/Fahrul-uzi/gin-full/config"
	"github.com/Fahrul-uzi/gin-full/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

// only admin to create new user
func createUser(c *gin.Context) {
	user := models.User{
		NIS:      c.PostForm("title"),
		Username: c.PostForm("desc"),
		Fullname: slug.Make(c.PostForm("title")),
		Password: c.PostForm("password"),
		Role:     c.PostForm("role"),
	}

	config.DB.Create(&user)
}
