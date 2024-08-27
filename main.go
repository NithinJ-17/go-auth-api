package main

import (
	"github.com/gin-gonic/gin"
	"go-auth-api/handlers"
	"go-auth-api/middlewares"
)

func main() {
	router := gin.Default()

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Protected route
	protected := router.Group("/protected")
	protected.Use(middlewares.JWTAuth())
	{
		protected.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "This is a protected route"})
		})
	}

	router.Run(":8000")
}
