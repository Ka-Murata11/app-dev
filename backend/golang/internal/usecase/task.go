package usecase

import (
	"myapp/infrastructure/entity"
	"myapp/internal/model"
	"myapp/internal/repository"
)

type TaskUsecase interface {
	CreateTask(task model.TaskRequest) error
	GetTask(task model.TaskID) (model.TaskResponse, error)
	UpdateTask(task model.TaskRequest) error
	DeleteTask(taskId model.TaskID) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo}
}

func (u *taskUsecase) CreateTask(task model.TaskRequest) error {
	taskEntity := entity.Task{
		TaskID: task.TaskID,
		Title:  task.Title,
	}

	if err := u.taskRepo.CreateTask(taskEntity); err != nil {
		return err
	}

	return nil
}

func (u *taskUsecase) GetTask(taskID model.TaskID) (model.TaskResponse, error) {
	taskEntity, err := u.taskRepo.GetTask(taskID.ID)
	if err != nil {
		return model.TaskResponse{}, err
	}

	taskRes := model.TaskResponse{
		TaskID:    taskEntity.TaskID,
		Title:     taskEntity.Title,
		CreatedAt: taskEntity.CreatedAt,
		UpdatedAt: taskEntity.UpdatedAt,
	}

	return taskRes, nil
}

func (u *taskUsecase) UpdateTask(taskReq model.TaskRequest) error {
	taskEntity, err := u.taskRepo.GetTask(taskReq.TaskID)
	if err != nil {
		return err
	}

	taskEntity.Title = taskReq.Title

	if err := u.taskRepo.UpdateTask(taskEntity); err != nil {
		return err
	}

	return nil

}

func (u *taskUsecase) DeleteTask(taskID model.TaskID) error {

	if err := u.taskRepo.DeleteTask(taskID.ID); err != nil {
		return err
	}

	return nil

}
