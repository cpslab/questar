package main

import (
	"github.com/cpslab/questar/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	users.UsersRegister(v1.Group("/users"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messages":    "pong",
			"testMessage": "ping",
		})
	})

	router.Run(":8000")
}
