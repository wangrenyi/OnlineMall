package repository

import (
	"github.com/jinzhu/gorm"
	"onlinemall/db"
	"onlinemall/logging"
	"reflect"
)

// https://github.com/Ksloveyuan/gorm-ex
// https://www.bbsmax.com/A/kjdwl7qBzN/

type BaseDAO struct {
	Connect *gorm.DB
}

func NewBaseDAO() BaseDAO {
	return BaseDAO{db.Connect()}
}

//func (baseDAO baseDAO) SaveTx(model interface{}) {
//	connect := baseDAO.Connect
//
//	tx := connect.Begin()
//	baseDAO.Save(model, tx)
//}

func (baseDAO BaseDAO) Save(model interface{}) error {
	verifyType(model)

	connect := baseDAO.Connect

	if err := connect.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (baseDAO BaseDAO) SaveAll(models ...interface{}) error {
	connect := baseDAO.Connect

	var err error
	for _, model := range models {
		if err := connect.Create(model).Error; err != nil {
			err = err
		}
	}

	return err
}

//func (baseDAO baseDAO) UpdateTx(model interface{}) {
//	connect := baseDAO.Connect
//
//	tx := connect.Begin()
//	baseDAO.Update(model, tx)
//}

func (baseDAO BaseDAO) Update(model interface{}) error {
	verifyType(model)
	connect := baseDAO.Connect

	modelValue := valueOf(model).Elem()
	if err := connect.Model(model).Updates(modelValue).Error; err != nil {
		return err
	}

	return nil
}

func (baseDAO BaseDAO) UpdateAll(models ...interface{}) error {

	connect := baseDAO.Connect

	var err error
	for _, model := range models {
		modelValue := valueOf(model).Elem()
		if err := connect.Model(model).Updates(modelValue).Error; err != nil {
			err = err
		}
	}

	return err
}

func (baseDAO BaseDAO) Delete(model interface{}) error {
	verifyType(model)

	connect := baseDAO.Connect
	if err := connect.Delete(model).Error; err != nil {
		return err
	}

	return nil
}

func (baseDAO BaseDAO) UniqueEntityById(model interface{}) {
	verifyType(model)

	connect := baseDAO.Connect

	_id := getFieldValue(model, "ID")
	connect.Where("id = ?", _id).First(model)
}

func (baseDAO BaseDAO) UniqueEntityByCondition(model interface{}, params map[string]interface{}) {
	verifyType(model)

	connect := baseDAO.Connect

	for k, v := range params {
		if k != "PageIndex" && k != "PageSize" && k != "OrderBy" {
			connect = whereEq(connect, k, v)
		}
	}
	connect.Find(model)
}

func (baseDAO BaseDAO) SelectAll(models interface{}) {
	verifyType(models)

	connect := baseDAO.Connect
	connect.Find(&models)
}

func (baseDAO BaseDAO) SelectEntityPaging(models interface{}, params map[string]interface{}) {
	verifyType(models)

	connect := baseDAO.Connect

	var pageIndex = params["PageIndex"].(int)
	var pageSize = params["PageSize"].(int)
	offset := 0
	if pageIndex > 0 {
		offset = pageIndex * pageSize
	}

	connect.Offset(offset).Limit(pageSize)
	for k, v := range params {
		if k != "PageIndex" && k != "PageSize" && k != "OrderBy" {
			connect = whereEq(connect, k, v)
		}
	}

	connect.Order(params["OrderBy"], true).Find(&models)
}

func whereEq(connect *gorm.DB, propertyName string, propertyValue interface{}) *gorm.DB {
	return connect.Where(propertyName+" = ?", propertyValue)
}

func (baseDAO BaseDAO) Count(model interface{}, params map[string]interface{}) int {
	verifyType(model)

	connect := baseDAO.Connect

	connect.Model(model)
	for k, v := range params {
		if k != "PageIndex" && k != "PageSize" && k != "OrderBy" {
			connect = whereEq(connect, k, v)
		}
	}

	count := 0
	connect.Count(&count)

	return count
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

func getFieldValue(model interface{}, fieldName string) interface{} {
	_value := valueOf(model).Elem()
	_type := _value.Type()

	for i := 0; i < _value.NumField(); i++ {
		field := _value.Field(i)
		if _type.Field(i).Name == fieldName {
			return field.Interface()
		}
	}

	return nil
}
