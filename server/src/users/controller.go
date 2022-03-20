package users

import (
	"fmt"
	"net/http"

	"github.com/cpslab/questar/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRegistration(c *gin.Context) {
	var db *gorm.DB = common.Init()
	var count int64
	db.Model(&User{}).Count(&count)
	if count == 3 {
		db.Select("UserName", "Password").Create(&User{UserName: "testuser4", Password: "test4"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func UserRetrieve(c *gin.Context) {
	var db *gorm.DB = common.Init()
	var users []User
	result := db.Find(&users)
	fmt.Println(result)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"results": result,
	})
}

func UserUpdate(c *gin.Context) {}

func UserDelete(c *gin.Context) {}
