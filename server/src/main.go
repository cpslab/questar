package main

import (
	"net/http"

	"github.com/cpslab/questar/users"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	router.GET("/testuser", users.UserRetrieve)
	router.GET("/testdb", func(c *gin.Context) {
		dsn := "host=localhost user=admin password=admin dbname=questar port=5432 sslmode=disable TimeZone=Asia/Tokyo"

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
				"error":   err,
			})
		}

		db.Select("UserName", "Password").Create(&users.User{UserName: "testuser4", Password: "test4"})
		var users []users.User
		// result := db.Find(&users)
		db.Find(&users)

		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
			// "result":  result,
		})
	})

	router.Run(":8000")
}
