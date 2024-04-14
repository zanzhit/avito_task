package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/errs"
)

func (h *Handler) getBanner(c *gin.Context) {
	var input banner.Banner
	var err error

	featureId := c.Query("feature_id")
	tagId := c.Query("tag_id")

	if featureId != "" {
		input.Feature, err = strconv.Atoi(featureId)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	if tagId != "" {
		tag, err := strconv.Atoi(tagId)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		input.Tag = append(input.Tag, tag)
	}

	limit := 0
	offset := 0

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	if c.Query("offset") != "" {
		offset, err = strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	banner, err := h.services.Banner.GetBanner(input, limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, banner)
}

func (h *Handler) createBanner(c *gin.Context) {
	var input banner.Banner
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Banner.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if id == -1 {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"banner_id": id,
	})
}

func (h *Handler) updateBanner(c *gin.Context) {
	bannerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "wrong banner id format")
		return
	}

	var input banner.UpdateBanner
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.Banner.Patch(input, bannerId)
	if err != nil {
		switch err.(type) {
		case *errs.ErrBannerNotFound:
			c.Status(http.StatusNotFound)
			return
		case *errs.ErrBannerNotUnique:
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteBanner(c *gin.Context) {
	bannerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "wrong banner id format")
		return
	}

	err = h.services.Banner.Delete(bannerId)
	if err != nil {
		switch err.(type) {
		case *errs.ErrBannerNotFound:
			c.Status(http.StatusNotFound)
			return
		default:
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Status(http.StatusNoContent)
}
