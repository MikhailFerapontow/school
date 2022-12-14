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

	signIn := router.Group("/signIn")
	{
		signIn.GET("", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		auth := router.Group("/auth")
		{
			auth.POST("/signUpAdmin", h.signUpAdmin)
			auth.POST("/signUpTeacher", h.signUpTeacher)
			auth.POST("/signUpStudent", h.signUpStudent)
		}

		classroom := api.Group("/class")
		{
			classroom.GET("", h.getClassroom)
		}

		guardians := api.Group("/guardian")
		{
			guardians.GET("", h.getAllGuardians)
			guardians.POST("", h.createGuardian)
		}
	}

	return router
}
