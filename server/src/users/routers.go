package users

import "github.com/gin-gonic/gin"

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/registration", UsersRegistration)
}

func UsersRegistration(c *gin.Context) {}
