package handler

import (
	"github.com/MikhailFerapontow/school/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	guardians := router.Group("/guardian")
	{
		guardians.GET("", h.getAllGuardians)
		guardians.POST("", h.createGuardian)
	}

	return router
}
