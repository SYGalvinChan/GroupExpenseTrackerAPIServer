package router

import (
	"GroupExpenseTracker/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	telegramAPI := r.Group("/api/v0/telegram")
	
	telegramAPI.POST("/newuser", handlers.NotImplementedYet)
	telegramAPI.POST("/newgroup", handlers.NotImplementedYet)
	telegramAPI.POST("/getusersingroup", handlers.NotImplementedYet)
	telegramAPI.POST("/getgroupswithuser", handlers.NotImplementedYet)
	telegramAPI.POST("/addusertogroup", handlers.NotImplementedYet)
	telegramAPI.POST("/addexpense", handlers.NotImplementedYet)
	telegramAPI.POST("/getexpenses", handlers.NotImplementedYet)
	telegramAPI.POST("/deleteexpense", handlers.NotImplementedYet)
	telegramAPI.POST("/editexpense", handlers.NotImplementedYet)
	telegramAPI.POST("/settleup", handlers.NotImplementedYet)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}