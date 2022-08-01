package initialize

import (
	"2022summer/global"
	"2022summer/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func InitMySQL() (err error) {
	// 配置文件
	viper.SetConfigFile("./config.yml") // 指定配置文件路径
	err = viper.ReadInConfig()          // 读取配置信息
	if err != nil {                     // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()

	// 配置数据
	var addr, port, user, password, dbstr string
	if viper.GetBool("debug") {
		addr = viper.GetString("db1.addr")
		port = viper.GetString("db1.port")
		user = viper.GetString("db1.user")
		password = viper.GetString("db1.password")
		dbstr = viper.GetString("db1.dbstr")
	} else {
		addr = viper.GetString("db2.addr")
		port = viper.GetString("db2.port")
		user = viper.GetString("db2.user")
		password = viper.GetString("db2.password")
		dbstr = viper.GetString("db2.dbstr")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, addr, port, dbstr)

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
