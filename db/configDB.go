package db

import (
	"onlinemall/config"
	"onlinemall/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var connect *gorm.DB

func init() {
	db, err := gorm.Open(config.DatasourceConfig.Dialect, config.DatasourceConfig.Dburl)
	if err != nil {
		panic("failed to connect database")
	}
	db.SetLogger(logging.DBLogger())
	db.DB().SetMaxIdleConns(config.DatasourceConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DatasourceConfig.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.DatasourceConfig.ConnMaxLifetime))
	//禁用表名复数,如果只是部分表需要使用源表名，请在实体类中声明TableName的构造函数
	//db.SingularTable(true)

	// 创建表时添加表后缀
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate((&model.User{}))
	connect = db
}

func Connect() *gorm.DB {
	return connect
}
