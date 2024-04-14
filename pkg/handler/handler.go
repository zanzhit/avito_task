package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zanzhit/avito_task/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	userBanner := router.Group("/user_banner", h.userIdentity)
	{
		userBanner.GET("", h.getUserBanner)
	}

	entities := router.Group("/entities", h.adminOnly)
	{
		banner := entities.Group("/banner")
		{
			banner.GET("", h.getBanner)
			banner.POST("", h.createBanner)
			banner.PATCH("/:id", h.updateBanner)
			banner.DELETE("/:id", h.deleteBanner)
		}

		feature := entities.Group("/feature")
		{
			feature.POST("", h.createFeature)
		}

		tag := entities.Group("/tag")
		{
			tag.POST("", h.createTag)
		}
	}

	return router
}
