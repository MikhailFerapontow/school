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

	auth := router.Group("/auth")
	{
		auth.POST("/signUpAdmin", h.signUpAdmin)
		auth.POST("/signUpTeacher", h.signUpTeacher)
		auth.POST("/signUpStudent", h.signUpStudent)
		auth.POST("/signIn", h.signIn)
	}

	guardians := router.Group("/guardian")
	{
		guardians.GET("", h.getAllGuardians)
		guardians.POST("", h.createGuardian)
	}

	return router
}
