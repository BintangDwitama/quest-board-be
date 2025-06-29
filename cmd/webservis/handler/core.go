package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/BintangDwitama/quest-board-be/internal/domain/constants"
	"github.com/BintangDwitama/quest-board-be/internal/domain/model"
)

type CreateTaskHandler func(*gin.Context, model.Task) error

func CreateTask(fn CreateTaskHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := fn(ctx, model.Task{})
		if err != nil {
			if errors.As(err, &constants.GeneralErrorConcrete) {
				ctx.AbortWithStatusJSON(
					constants.GeneralErrorConcrete.Code,
					constants.FailedResponse(
						constants.GeneralErrorConcrete.Message, nil,
					),
				)
				return
			} else {
				ctx.AbortWithStatusJSON(
					500,
					constants.FailedResponse(
						"Unknown Error", nil,
					),
				)
			}
		}

		ctx.JSON(
			200,
			constants.SuccessResponse("Request Success", nil),
		)
	}
}
