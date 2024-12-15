package model

import "time"

type TaskID struct {
	ID string `query:"task_id"`
}

type TaskRequest struct {
	TaskID string `json:"task_id"`
	Title  string `json:"title"`
}

type TaskResponse struct {
	TaskID    string    `json:"task_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}
