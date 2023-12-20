package controller

import (
	"changeme/Dao"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(db *gorm.DB) {
	DB = db

	// 自动迁移
	db.AutoMigrate(&Dao.CountDownDao{})
	db.AutoMigrate(&Dao.UserDao{})
}
