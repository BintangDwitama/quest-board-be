package router

import (
	"github.com/gin-gonic/gin"

	"github.com/BintangDwitama/quest-board-be/cmd/webservis/handler"
	"github.com/BintangDwitama/quest-board-be/cmd/webservis/middleware"
	"github.com/BintangDwitama/quest-board-be/internal/domain/usecase"
)

func CoreRouter(router *gin.Engine, usecase usecase.Usecase) {
	m := middleware.NewMiddleware(usecase)
	publicSecure := middleware.NewMiddlewareFactory()
	publicSecure.Use(m.Logger())

	router.GET("/tasks")
	router.POST("/tasks", append(publicSecure.Middlewares(), handler.CreateTask(usecase.CreateNewTask))...)
	router.PUT("/tasks")
	router.DELETE("/tasks")
}
