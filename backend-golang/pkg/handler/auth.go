package handler

import (
	"errors"
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
		return
	}

	err := h.services.Authorization.RegisterStudent(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created Student",
	})
}

func (h *Handler) signUpTeacher(c *gin.Context) {
	var input school.RegisterTeacher

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.RegisterTeacher(input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created Teacher",
	})
}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (h *Handler) getUserRole(c *gin.Context) (string, error) {
	login, ok := c.Get(userCtx)
	if !ok {
		return "", errors.New("empty context")
	}

	role, err := h.services.Authorization.GetUserRole(login)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (h *Handler) CheckRole(c *gin.Context, roles ...string) error {
	role, err := h.getUserRole(c)

	if err != nil {
		return err
	}
	var flag bool

	for _, elem := range roles {
		if role == elem {
			flag = true
		}
	}

	if !flag {
		return errors.New("Unauthorized")
	}

	return nil
}
