package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.Status(http.StatusUnauthorized)
		return
	}

	if len(headerParts[1]) == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	_, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) adminOnly(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.Status(http.StatusUnauthorized)
		return
	}

	if len(headerParts[1]) == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if role != "admin" {
		c.Status(http.StatusForbidden)
		return
	}

	c.Next()
}
