package handler

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/constants"
	"github.com/gin-gonic/gin"
)

func GetTaskList(ctx *gin.Context) {
	ctx.JSON(
		503,
		constants.FailedResponse("Service Unavailable", nil),
	)
}

func CreateTask(ctx *gin.Context) {
	ctx.JSON(
		503,
		constants.FailedResponse("Service Unavailable", nil),
	)
}

func UpdateTask(ctx *gin.Context) {
	ctx.JSON(
		503,
		constants.FailedResponse("Service Unavailable", nil),
	)
}

func DeleteTask(ctx *gin.Context) {
	ctx.JSON(
		503,
		constants.FailedResponse("Service Unavailable", nil),
	)
}