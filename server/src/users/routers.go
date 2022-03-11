package users

import "github.com/gin-gonic/gin"

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/user/register", UserRegister)
	router.GET("/user", UserRetrieve)
	router.PUT("/user/update", UserUpdate)
	router.DELETE("/user/delete", UserDelete)
}
