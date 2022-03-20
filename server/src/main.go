package main

import (
	"net/http"

	"github.com/cpslab/questar/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")
	users.UserRegister(v1.Group("/user"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages":    "pong",
			"testMessage": "ping",
		})
	})

	router.Run(":8000")
}
