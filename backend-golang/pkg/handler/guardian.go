package handler

import (
	"net/http"

	"github.com/MikhailFerapontow/school"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createGuardian(c *gin.Context) {

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
