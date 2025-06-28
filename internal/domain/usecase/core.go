package usecase

import (
	"github.com/BintangDwitama/quest-board-be/internal/domain/constants"
	"github.com/BintangDwitama/quest-board-be/internal/domain/model"
)

type CoreUsecase interface {
	CreateNewTask(newTask model.Task) error
	GetAllTasks(userId int64) ([]model.Task, error)
	UpdateTask(taskId int64, taskData model.Task) error
	DeleteTask(taskId int64) error
}

type coreUsecase struct{}

func NewCoreUsecase() *coreUsecase {
	return &coreUsecase{}
}

func (core *coreUsecase) CreateNewTask(newTask model.Task) error {
	return constants.NewGeneralError(503, "Service no available")
}

func (core *coreUsecase) GetAllTasks(userId int64) ([]model.Task, error) {
	return nil, constants.NewGeneralError(503, "Service no available")
}

func (core *coreUsecase) UpdateTask(taskId int64, taskData model.Task) error {
	return constants.NewGeneralError(503, "Service no available")
}

func (core *coreUsecase) DeleteTask(taskId int64) error {
	return constants.NewGeneralError(503, "Service no available")
}
