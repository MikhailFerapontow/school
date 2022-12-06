package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getClassroom(c *gin.Context) {
	login, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"login": login,
	})
}
