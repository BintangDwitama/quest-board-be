package transport

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/handler"
	"github.com/gin-gonic/gin"
)

func StartHTTPServer() {
	router := gin.Default()

	router.GET("/ping", handler.Ping)

	router.Run()
}