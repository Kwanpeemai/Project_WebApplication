package main

import (
	"github.com/gin-gonic/gin"
	"go.mod/controllers"
)

func main() {
	router := gin.Default()

	controllers.ConnectDatabase()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
