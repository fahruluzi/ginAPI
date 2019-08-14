package main

import (
	"fmt"
	"os"

	"github.com/Fahrul-uzi/gin-full/config"
	"github.com/Fahrul-uzi/gin-full/routes"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	// Inisialisasi Database
	config.InitDB()
	defer config.DB.Close()

	gotenv.Load()

	// Router
	router := gin.Default()

	fmt.Println(os.Getenv("CLIENT_ID_GITHUB"))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/auth/:provider", routes.RedirectHandler)
		v1.GET("/auth/:provider/callback", routes.CallbackHandler)
		articles := v1.Group("/article")
		{
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", routes.PostArticle)
		}
	}

	router.Run()
}
