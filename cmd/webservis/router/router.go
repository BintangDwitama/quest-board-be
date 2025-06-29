package router

import (
	"github.com/gin-gonic/gin"

	"github.com/BintangDwitama/quest-board-be/internal/domain/usecase"
)

func Init(u usecase.Usecase) *gin.Engine {
	router := gin.Default()
	CoreRouter(router, u)
	return router
}
