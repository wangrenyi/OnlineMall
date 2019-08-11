package repository

import (
	"github.com/jinzhu/gorm"
	"onlinemall/db"
	"onlinemall/logging"
	"reflect"
)

type BaseDAO struct {
	Connect *gorm.DB
}

func NewBaseDAO() *BaseDAO {
	return &BaseDAO{db.Connect()}
}

func (baseDAO *BaseDAO) SaveTx(model interface{}) {
	connect := baseDAO.Connect

	tx := connect.Begin()
	baseDAO.Save(model, tx)
}

func (baseDAO *BaseDAO) Save(model interface{}, tx *gorm.DB) {
	verifyType(model)

	defer commitTrans(tx)

	if err := tx.Create(model).Error; err != nil {
		panicError(err)
	}
}

func commitTrans(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
}

func panicError(err interface{}) {
	logging.Info(err)
	panic(err)
}

//判断传入的model是否是指针类型
func verifyType(model interface{}) {
	if !(valueOf(model).Kind() == reflect.Ptr) {
		panicError("the entity must be a pointer type.")
	}
}

func valueOf(model interface{}) reflect.Value {
	return reflect.ValueOf(model)
}
