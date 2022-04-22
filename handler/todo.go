package handler

import (
	"go-todo-service/router"
	"go-todo-service/service"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todoList)

}

func (t todoHandler) GetByID(c *router.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	todo, err := t.todoSrv.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (t todoHandler) NewTodo(c *router.Context) {
	input := service.TodoRequest{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	_, err = t.todoSrv.NewTodo(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (t todoHandler) UpdateByID(c *router.Context) {
	input := service.TodoRequest{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = t.todoSrv.UpdateTodo(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
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
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = t.todoSrv.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
