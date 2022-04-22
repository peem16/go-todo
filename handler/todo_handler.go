package handler

type TodoRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TodoResponse struct {
	TodoID    int    `json:"todoID"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type TodoService interface {
	NewTodo(TodoRequest) (*TodoResponse, error)
	GetTodo(int) (*TodoResponse, error)
	GetTodos() ([]TodoResponse, error)
	UpdateTodo(int) error
	DeleteTodo(int) error
}
