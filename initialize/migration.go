package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orderManagement/models"
)

func Migrate() {
	dsn := "root:password@tcp(127.0.0.1:3306)/sellerapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		panic(err)
	}
}
