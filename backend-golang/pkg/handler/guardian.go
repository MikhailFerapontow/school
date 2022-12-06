package handler

import (
	"net/http"

	"github.com/MikhailFerapontow/school"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createGuardian(c *gin.Context) {

	err := h.CheckRole(c, "student")
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input school.GuardianInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Guardian.CreateGuardian(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created Guardian",
	})
}

type getAllGuardiansResponse struct {
	Data []school.Guardian `json:"data"`
}

func (h *Handler) getAllGuardians(c *gin.Context) {

	guardians, err := h.services.Guardian.GetAll()
	if err != nil {
		NewErrorResponse(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusFound, getAllGuardiansResponse{
		Data: guardians,
	})
}
