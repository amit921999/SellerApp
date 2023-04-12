package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var dsn = "root:password@tcp(127.0.0.1:3306)/sellerapp?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
