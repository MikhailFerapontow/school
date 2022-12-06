package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getClassroom(c *gin.Context) {
	role, err := h.getUserRole(c)

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if role != "student" {
		NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"role": role,
	})
}
