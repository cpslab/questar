package common

import (
	"fmt"

	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	db, err := gorm.Open("postgres", "")
	if err != nil {
		fmt.Println("error")
	}

	Db = db
	return Db
}
