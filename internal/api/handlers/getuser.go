package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"web3pay/internal/api/paths"
)

func (h *Handlers) GetUser(c *gin.Context) {
	var userId *paths.UserId
	err := c.ShouldBindUri(&userId)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	user, err := h.db.GetUser(c.Request.Context(), userId.Id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &user)
}
