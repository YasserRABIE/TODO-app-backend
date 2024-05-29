package main

import (
	"os"

	"github.com/YasserRABIE/authentication-porject/handlers"
	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/api/register", handlers.RegisterHandler, middleware.HandleAuth)
	r.POST("/api/login", handlers.LoginHandler, middleware.HandleAuth)
	r.GET("/api/account", middleware.RequireAuth, handlers.GetAccountHandler)

	r.Run(os.Getenv("PORT"))
}
