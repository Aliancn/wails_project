package Dao

import "gorm.io/gorm"

type TodoListDao struct {
	gorm.Model
	ID       uint
	TaskName string
	Competed bool
}

func (t *TodoListDao) GetAll(db *gorm.DB) ([]TodoListDao, error) {
	var tmp []TodoListDao
	err := db.Find(&tmp).Error
	return tmp, err
}

func (t *TodoListDao) TableName() string {
	return "todo_list"
}

func (t *TodoListDao) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *TodoListDao) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *TodoListDao) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}
