package service

import (
	"go-todo-service/errs"
	"go-todo-service/logs"
	"go-todo-service/repository"
)

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return todoService{todoRepo: todoRepo}
}

func (t todoService) NewTodo(reqTodo TodoRequest) (*TodoResponse, error) {
	count, err := t.todoRepo.Count()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	repoTodo := repository.Todo{
		TodoID: count + 1,
		Title:  reqTodo.Title,
		Status: reqTodo.Status,
	}

	err = t.todoRepo.Create(repoTodo)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	return nil, nil
}
func (t todoService) GetTodo(id int) (*TodoResponse, error) {
	todo, err := t.todoRepo.GetByID(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	todoResponse := TodoResponse{
		TodoID:    int(id),
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return &todoResponse, nil
}
func (t todoService) GetTodos() ([]TodoResponse, error) {
	todoList, err := t.todoRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	todoRes := []TodoResponse{}

	for _, todo := range todoList {
		todoRes = append(todoRes, TodoResponse{
			int(todo.TodoID),
			todo.Title,
			todo.Status,
			todo.CreatedAt,
			todo.UpdatedAt,
		})
	}

	return todoRes, nil
}
func (t todoService) UpdateTodo(id int, reqTodo TodoRequest) error {
	repoTodo := repository.Todo{
		Title:  reqTodo.Title,
		Status: reqTodo.Status,
	}
	err := t.todoRepo.UpdateByID(id, repoTodo)

	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t todoService) DeleteTodo(id int) error {
	err := t.todoRepo.Delete(id)

	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}
