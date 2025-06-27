package transport

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/handler"
	"github.com/gin-gonic/gin"
)

func StartHTTPServer() {
	router := gin.Default()

	router.GET("/ping", handler.Ping)

	router.GET("/tasks", handler.GetTaskList)
	router.POST("/tasks", handler.CreateTask)
	router.PUT("/tasks", handler.UpdateTask)
	router.DELETE("/tasks", handler.DeleteTask)

	router.Run()
}