package interfaces

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserAccount(c *gin.Context)
	GetUserBalance(c *gin.Context)
	DeleteUser(c *gin.Context)
}
