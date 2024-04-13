package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	banner "github.com/zanzhit/avito_task"
)

func (h *Handler) getUserBanner(c *gin.Context) {
	var input banner.UserBanner
	var err error

	featureId := c.Query("feature_id")
	tagId := c.Query("tag_id")

	if featureId == "" && tagId == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid feature/tag param")
	}

	input.Feature, err = strconv.Atoi(featureId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid feature param")
		return
	}

	input.Tag, err = strconv.Atoi(tagId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid tag param")
		return
	}

	lastRevision := false

	if c.Query("use_last_revision") != "" {
		lastRevision, err = strconv.ParseBool(c.Query("use_last_revision"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid revision param")
			return
		}
	}

	banner, err := h.services.UserBanner.GetUserBanner(input, lastRevision)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, banner)
}
