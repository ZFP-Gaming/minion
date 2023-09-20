package main

import (
	"github.com/ZFP-Gaming/minion/initializers"
	"github.com/ZFP-Gaming/minion/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDb()
}

func main() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	r := gin.Default()
	r.Use(cors.New(config))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/search", routes.FindUserByName)
	r.GET("/all", routes.AllUsers)
	r.Run(":8080")
}
