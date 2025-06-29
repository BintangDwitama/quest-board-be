package usecase

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/constants"
	"github.com/BintangDwitama/quest-board-be/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type Usecase interface {
	CreateNewTask(ctx *gin.Context, newTask model.Task) error
	GetAllTasks(ctx *gin.Context, userId int64) ([]model.Task, error)
	UpdateTask(ctx *gin.Context, taskId int64, taskData model.Task) error
	DeleteTask(ctx *gin.Context, taskId int64) error
}

type usecase struct{}

func NewUsecase() Usecase {
	return &usecase{}
}

func (core *usecase) CreateNewTask(ctx *gin.Context, newTask model.Task) error {
	return constants.NewGeneralError(503, "Service no available")
}

func (core *usecase) GetAllTasks(ctx *gin.Context, userId int64) ([]model.Task, error) {
	return nil, constants.NewGeneralError(503, "Service no available")
}

func (core *usecase) UpdateTask(ctx *gin.Context, taskId int64, taskData model.Task) error {
	return constants.NewGeneralError(503, "Service no available")
}

func (core *usecase) DeleteTask(ctx *gin.Context, taskId int64) error {
	return constants.NewGeneralError(503, "Service no available")
}
