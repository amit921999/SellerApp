package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orderManagement/models"
)

func Migrate() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		panic(err)
	}

	sqlConn, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlConn.Close()
	if err != nil {
		panic(err)
	}
}
