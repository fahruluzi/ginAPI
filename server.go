package main

import (
	"github.com/Fahrul-uzi/gin-full/config"
	"github.com/Fahrul-uzi/gin-full/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi Database
	config.InitDB()
	defer config.DB.Close()

	// Router
	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		articles := v1.Group("/article/")
		{
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", routes.PostArticle)
		}
	}

	router.Run()
}
