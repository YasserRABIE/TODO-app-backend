package main

import (
	"os"

	"github.com/YasserRABIE/authentication-porject/handlers"
	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	r.POST("/register", handlers.HandleRegister)

	r.Run(os.Getenv("PORT"))
}
