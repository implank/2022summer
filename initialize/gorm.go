package initialize

import (
	"2022summer/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySQL() (err error) {
	// 配置数据
	addr := "43.138.77.133"
	port := "3306"
	user := "remote"
	password := "remote"
	dbstr := "2022summer"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, addr, port, dbstr)

	// 连接数据库
	global.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 迁移
	global.DB.AutoMigrate()

	return global.DB.DB().Ping()
}

func Close() {
	err := global.DB.Close()
	if err != nil {
		return
	}
}
