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

	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler, middleware.HandleAuth)

	r.Run(os.Getenv("PORT"))
}
