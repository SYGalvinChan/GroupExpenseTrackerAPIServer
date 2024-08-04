package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotImplementedYet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"msg": "not implemented yet",
	})
}
