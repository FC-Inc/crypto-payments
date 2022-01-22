package handlers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"web3pay/internal/api/paths"
)

func (h *Handlers) CreateUser(c *gin.Context) {
	var userId *paths.UserId
	err := c.ShouldBindUri(&userId)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = c.ShouldBindQuery("account")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	account := c.Query("account")
	if !common.IsHexAddress(account) {
		c.Status(http.StatusBadRequest)
		return
	}
	err = h.db.CreateUser(c.Request.Context(), userId.Id, account)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusCreated)
}
