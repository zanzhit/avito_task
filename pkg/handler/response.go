package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	badRequest = "Некорректные данные"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
