package main

import Gin "github.com/gin-gonic/gin"

func main() {
	var gin = Gin.Default()

	gin.GET("/api/v1", func(context *Gin.Context) {
		context.JSON(200, Gin.H{
			"message": "Hello Gin",
		})
	})

	gin.Run()
}
