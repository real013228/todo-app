package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/real013228/todo-app"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}
