package main

import (
	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Run()
}
