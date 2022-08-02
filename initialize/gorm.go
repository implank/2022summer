package initialize

import (
	"2022summer/global"
	"2022summer/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySQL() (err error) {
	// 配置数据
	var addr, port, user, password, dbstr string
	if global.VP.GetBool("debug") {
		addr = global.VP.GetString("db1.addr")
		port = global.VP.GetString("db1.port")
		user = global.VP.GetString("db1.user")
		password = global.VP.GetString("db1.password")
		dbstr = global.VP.GetString("db1.dbname")
	} else {
		addr = global.VP.GetString("db2.addr")
		port = global.VP.GetString("db2.port")
		user = global.VP.GetString("db2.user")
		password = global.VP.GetString("db2.password")
		dbstr = global.VP.GetString("db2.dbname")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, addr, port, dbstr)
	fmt.Printf("dsn: %s\n", dsn)
	// 连接数据库
	global.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 迁移
	global.DB.AutoMigrate(
		&model.User{},
		&model.Identity{},
		&model.Group{},
		&model.Proj{},
		&model.Prototype{},
		&model.Uml{},
		&model.Document{},
	)

	return global.DB.DB().Ping()
}

func Close() {
	err := global.DB.Close()
	if err != nil {
		return
	}
}
