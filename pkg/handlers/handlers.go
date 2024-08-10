package handlers

import (
	"GroupExpenseTracker/pkg/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotImplementedYet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"msg": "not implemented yet",
	})
}

func NewUser(c *gin.Context) {
	storage.NewUser()
}
