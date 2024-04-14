package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	banner "github.com/zanzhit/avito_task"
)

func (h *Handler) createFeature(c *gin.Context) {
	var input banner.Feature
	var err error

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.services.Feature.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
