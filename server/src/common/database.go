package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=admin password=admin dbname=questar port=db sslmode=disable TimeZone=Asia/Tokyo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error(Init): ", err)
	}

	Db = db
	return Db
}
