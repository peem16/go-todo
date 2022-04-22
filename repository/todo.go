package repository

import "time"

type Todo struct {
	TodoID    int    `db:"todoid"`
	Title     string `db:"title"`
	Status    string `db:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoRepository interface {
	Create(Todo) error
	GetAll() ([]Todo, error)
	GetByID(int) (*Todo, error)
	UpdateByID(int, Todo) error
	Delete(int) error
}
