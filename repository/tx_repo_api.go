package repository

import "github.com/jinzhu/gorm"

type TxOperater interface {
	Save(tx *gorm.DB, model interface{}) error

	SaveAll(tx *gorm.DB, models ...interface{}) error

	Update(tx *gorm.DB, model interface{}) error

	UpdateAll(tx *gorm.DB, models ...interface{}) error

	Delete(tx *gorm.DB, model interface{}) error

	DeleteById(tx *gorm.DB, id interface{}) error
}
