package transport

import "github.com/gin-gonic/gin"

func StartHTTPServer() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Service Available",
		})
	})

	router.Run()
}