package transport

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/handler"
	"github.com/BintangDwitama/quest-board-be/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

func StartHTTPServer() {
	router := gin.Default()

	coreUsecase := usecase.NewCoreUsecase()
	coreHandler := handler.NewCoreHandler(coreUsecase)

	router.GET("/ping", handler.Ping)

	router.GET("/tasks", coreHandler.GetTaskList)
	router.POST("/tasks", coreHandler.CreateTask)
	router.PUT("/tasks", coreHandler.UpdateTask)
	router.DELETE("/tasks", coreHandler.DeleteTask)

	router.Run()
}
