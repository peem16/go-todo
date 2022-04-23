package handler

import (
	"go-todo-service/errs"
	"go-todo-service/router"
	"go-todo-service/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type todoHandler struct {
	todoSrv service.TodoService
}

func NewTodoHandler(todoSrv service.TodoService) todoHandler {
	return todoHandler{todoSrv: todoSrv}
}

func (t todoHandler) GetAll(c *router.Context) {
	todoList, err := t.todoSrv.GetTodos()
	if err != nil {
		handleError(c, errs.NewBadRequest(err.Error()))
		return
	}

	c.JSON(http.StatusOK, todoList)

}

func (t todoHandler) GetByID(c *router.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, err)
		return
	}

	todo, err := t.todoSrv.GetTodo(id)
	if err != nil {
		handleError(c, errs.NewBadRequest(err.Error()))
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (t todoHandler) NewTodo(c *router.Context) {
	input := service.TodoRequest{}

	if err := c.ShouldBindJSON(&input); err == nil {
		validate := validator.New()
		if err := validate.Struct(&input); err != nil {
			handleError(c, errs.NewValidationError(err.Error()))
			c.Abort()
			return
		}
	}

	_, err := t.todoSrv.NewTodo(input)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (t todoHandler) UpdateByID(c *router.Context) {
	input := service.TodoRequest{}
	if err := c.ShouldBindJSON(&input); err == nil {
		validate := validator.New()
		if err := validate.Struct(&input); err != nil {
			handleError(c, errs.NewValidationError(err.Error()))
			c.Abort()
			return
		}
	}

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, err)
		return
	}

	err = t.todoSrv.UpdateTodo(id, input)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (t todoHandler) Delete(c *router.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, err)
		return
	}

	err = t.todoSrv.DeleteTodo(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
