package handler_dep

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type coreHandler struct {
	usecase usecase.Usecase
}

func NewCoreHandler(usecase usecase.Usecase) *coreHandler {
	return &coreHandler{
		usecase: usecase,
	}
}

func (core *coreHandler) GetTaskList(ctx *gin.Context) {
	// resp, err := core.usecase.GetAllTasks(0)
	// if err != nil {
	// 	if errors.As(err, &constants.GeneralErrorConcrete) {
	// 		ctx.AbortWithStatusJSON(
	// 			constants.GeneralErrorConcrete.Code,
	// 			constants.FailedResponse(
	// 				constants.GeneralErrorConcrete.Message, nil,
	// 			),
	// 		)
	// 		return
	// 	} else {
	// 		ctx.AbortWithStatusJSON(
	// 			500,
	// 			constants.FailedResponse(
	// 				"Unknown Error", nil,
	// 			),
	// 		)
	// 	}
	// }

	// ctx.JSON(
	// 	200,
	// 	constants.SuccessResponse("Request Success", resp),
	// )
}

func (core *coreHandler) UpdateTask(ctx *gin.Context) {
	// err := core.usecase.UpdateTask(0, model.Task{})
	// if err != nil {
	// 	if errors.As(err, &constants.GeneralErrorConcrete) {
	// 		ctx.AbortWithStatusJSON(
	// 			constants.GeneralErrorConcrete.Code,
	// 			constants.FailedResponse(
	// 				constants.GeneralErrorConcrete.Message, nil,
	// 			),
	// 		)
	// 		return
	// 	} else {
	// 		ctx.AbortWithStatusJSON(
	// 			500,
	// 			constants.FailedResponse(
	// 				"Unknown Error", nil,
	// 			),
	// 		)
	// 	}
	// }

	// ctx.JSON(
	// 	200,
	// 	constants.SuccessResponse("Request Success", nil),
	// )
}

func (core *coreHandler) DeleteTask(ctx *gin.Context) {
	// err := core.usecase.DeleteTask(0)
	// if err != nil {
	// 	if errors.As(err, &constants.GeneralErrorConcrete) {
	// 		ctx.AbortWithStatusJSON(
	// 			constants.GeneralErrorConcrete.Code,
	// 			constants.FailedResponse(
	// 				constants.GeneralErrorConcrete.Message, nil,
	// 			),
	// 		)
	// 		return
	// 	} else {
	// 		ctx.AbortWithStatusJSON(
	// 			500,
	// 			constants.FailedResponse(
	// 				"Unknown Error", nil,
	// 			),
	// 		)
	// 	}
	// }

	// ctx.JSON(
	// 	200,
	// 	constants.SuccessResponse("Request Success", nil),
	// )
}
