package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	// Setup database
	dsn := "mysql://root:root@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Products{})

	return db
}
