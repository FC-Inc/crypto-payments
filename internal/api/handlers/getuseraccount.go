package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) GetUserAccount(c *gin.Context) {
	err := c.ShouldBindUri("userId")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}
