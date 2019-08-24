package repository

type Operater interface {

	Save(model interface{}) error

	SaveAll(models ...interface{}) error

	Update(model interface{}) error

	UpdateAll(models ...interface{}) error

	Delete(model interface{}) error

	UniqueEntityById(model interface{}) interface{}

	UniqueEntityByCondition(model interface{}, params map[string]interface{})

	SelectAll(models interface{})

	SelectEntityPaging(models interface{}, params map[string]interface{})

	Count(model interface{}, params map[string]interface{}) int
}
