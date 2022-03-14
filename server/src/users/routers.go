package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/registration", UserRegistration)
	router.GET("/", UserRetrieve)
	router.PUT("/update", UserUpdate)
	router.DELETE("/delete", UserDelete)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages":    "pong",
			"testMessage": "ping",
		})
	})
}
