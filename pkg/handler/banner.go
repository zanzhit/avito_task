package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	banner "github.com/zanzhit/avito_task"
)

func (h *Handler) getBanner(c *gin.Context) {
	var input banner.Banner
	var err error

	featureId := c.Query("feature_id")
	tagId := c.Query("tag_id")

	if featureId != "" {
		input.Feature, err = strconv.Atoi(featureId)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid feature param")
			return
		}
	}

	if tagId != "" {
		tag, err := strconv.Atoi(tagId)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid tag param")
			return
		}

		input.Tag = append(input.Tag, tag)
	}

	limit := 0
	offset := 0

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid limit param")
			return
		}
	}

	if c.Query("offset") != "" {
		offset, err = strconv.Atoi(c.Query("offset"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid offset param")
			return
		}
	}

	banner, err := h.services.Banner.GetBanner(input, limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, []interface{}{
		banner,
	})
}

func (h *Handler) createBanner(c *gin.Context) {
	var input banner.Banner
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "")
		return
	}

	id, err := h.services.Banner.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateBanner(c *gin.Context) {

}

func (h *Handler) deleteBanner(c *gin.Context) {

}
