package service

import "time"

type TodoRequest struct {
	Title  string `json:"title" validate:"required"`
	Status string `json:"status"`
}

type TodoResponse struct {
	TodoID    int       `json:"todoID"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type TodoService interface {
	NewTodo(TodoRequest) (*TodoResponse, error)
	GetTodo(int) (*TodoResponse, error)
	GetTodos() ([]TodoResponse, error)
	UpdateTodo(int, TodoRequest) error
	DeleteTodo(int) error
}
