package Dao

import (
	"gorm.io/gorm"
)

type UserDao struct {
	gorm.Model
	Username    string `gorm:"type:varchar(20);not null"`
	Password    string `gorm:"type:varchar(20);not null"`
	Email       string `gorm:"type:varchar(20);not null"`
	Avatar      string `gorm:"type:varchar(60);not null"`
	Description string `gorm:"type:varchar(150);not null"`
}

func (userinfo *UserDao) CheckLogin(db *gorm.DB) (bool, *UserDao) {
	var tmp []UserDao
	err := db.Where("username = ? and password = ?", userinfo.Username, userinfo.Password).Find(&tmp).Error
	if err != nil || len(tmp) != 1 {
		return false, nil
	}
	return true, &tmp[0]
}

func (userinfo *UserDao) Register(db *gorm.DB) bool {
	err := db.Create(userinfo).Error
	if err != nil {
		return false
	}
	return true
}
