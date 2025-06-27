package handler

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/constants"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(
		200,
		constants.SuccessResponse("Service Available", nil),
	)
}