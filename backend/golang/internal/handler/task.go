package handler

import (
	"myapp/internal/model"
	"myapp/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TaskHandler interface {
	CreateTask(c echo.Context) error
	GetTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase}
}

func (h *taskHandler) CreateTask(c echo.Context) error {
	var taskRequest model.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return echo.ErrBadRequest
	}
	if err := c.Validate(&taskRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.taskUsecase.CreateTask(taskRequest); err != nil {
		return errToHTTPStatus(err)

	}
	return c.JSON(http.StatusOK, "success task create")
}

func (h *taskHandler) GetTask(c echo.Context) error {
	var taskID model.TaskID
	if err := c.Bind(&taskID); err != nil {
		return echo.ErrBadRequest
	}
	if err := c.Validate(&taskID); err != nil {
		return echo.ErrBadRequest
	}

	task, err := h.taskUsecase.GetTask(taskID)
	if err != nil {
		return errToHTTPStatus(err)

	}
	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	var taskRequest model.TaskRequest
	if err := c.Bind(&taskRequest); err != nil {
		return echo.ErrBadRequest
	}
	if err := c.Validate(&taskRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.taskUsecase.UpdateTask(taskRequest); err != nil {
		return errToHTTPStatus(err)

	}
	return c.JSON(http.StatusOK, "success task update")
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	var taskID model.TaskID
	if err := c.Bind(&taskID); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&taskID); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.taskUsecase.DeleteTask(taskID); err != nil {
		return errToHTTPStatus(err)

	}
	return c.JSON(http.StatusOK, "success task delete")
}

func errToHTTPStatus(err error) error {
	if err == gorm.ErrRecordNotFound {
		return echo.ErrBadRequest
	} else {
		return echo.ErrInternalServerError
	}
}
