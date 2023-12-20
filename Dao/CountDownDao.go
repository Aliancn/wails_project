package Dao

import (
	"gorm.io/gorm"
	"time"
)

type CountDownDao struct {
	gorm.Model
	Content  string
	DeadTime time.Time
	Active   bool
	Title    string
}

func (c *CountDownDao) GetAll(db *gorm.DB) ([]CountDownDao, error) {
	var tmp []CountDownDao
	err := db.Find(&tmp).Error
	return tmp, err
}

func (c *CountDownDao) TableName() string {
	return "count_down"
}

func (c *CountDownDao) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func (c *CountDownDao) Update(db *gorm.DB) error {
	return db.Save(c).Error
}

func (c *CountDownDao) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}

func (c *CountDownDao) Get(db *gorm.DB) error {
	return db.First(c).Error
}
