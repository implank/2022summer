package initialize

import (
	"2022summer/global"
	"2022summer/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySQL() {
	// 配置数据
	var addr, port, user, password, dbname string
	if global.VP.GetBool("debug") {
		addr = global.VP.GetString("db1.addr")
		port = global.VP.GetString("db1.port")
		user = global.VP.GetString("db1.user")
		password = global.VP.GetString("db1.password")
		dbname = global.VP.GetString("db1.dbname")
	} else {
		addr = global.VP.GetString("db2.addr")
		port = global.VP.GetString("db2.port")
		user = global.VP.GetString("db2.user")
		password = global.VP.GetString("db2.password")
		dbname = global.VP.GetString("db2.dbname")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, addr, port, dbname)

	// 连接数据库
	var err error
	global.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("数据库出问题啦: %s \n", err))
		return
	}

	// 迁移
	global.DB.AutoMigrate(
		//base
		&model.User{},
		//group
		&model.Identity{},
		&model.Group{},
		//project
		&model.Proj{},
		&model.Prototype{},
		&model.Uml{},
		&model.Document{},
	)

	// 检查数据库连接是否存在, 好像没啥用
	err = global.DB.DB().Ping()
	if err != nil {
		panic(fmt.Errorf("数据库出问题啦: %s \n", err))
		return
	}

	return
}

func Close() {
	err := global.DB.Close()
	if err != nil {
		panic(fmt.Errorf("数据库出问题啦: %s \n", err))
		return
	}
}
