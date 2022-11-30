package handler

import (
	"net/http"

	"github.com/MikhailFerapontow/school"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUpAdmin(c *gin.Context) {
	var input school.RegisterAdmin

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.RegisterAdmin(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created Admin",
	})
}

func (h *Handler) signUpStudent(c *gin.Context) {
	var input school.RegisterStudent

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.services.RegisterStudent(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created Student",
	})
}

func (h *Handler) signUpTeacher(c *gin.Context) {
	var input school.RegisterTeacher

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.services.RegisterTeacher(input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created Teacher",
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
