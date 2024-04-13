package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	banner "github.com/zanzhit/avito_task"
)

func (h *Handler) createTag(c *gin.Context) {
	var input banner.Tag
	var err error

	input.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	_, err = h.services.Tag.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}