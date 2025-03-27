package models

type Task struct {
	Id     int    `json:"id"`
	Title  string `json:"title" validate:"required"`
	Status string `json:"status" validate:"required"`
}
