package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=admin password=admin dbname=questar port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error(Init): ", err)
	}

	return db
}
