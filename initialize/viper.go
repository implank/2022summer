package initialize

import (
	"2022summer/global"
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() (err error) {
	// 配置文件
	v := viper.New()
	v.SetConfigFile("./config.yml") // 指定配置文件路径
	err = v.ReadInConfig()          // 读取配置信息
	if err != nil {                 // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	global.VP = v
	return err
}
