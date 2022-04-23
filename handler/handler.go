package handler

import (
	"go-todo-service/errs"
	"go-todo-service/router"
	"net/http"
)

func handleError(c *router.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.JSON(e.Code, map[string]interface{}{
			"error": err.Error(),
		})
	case error:
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
}
