package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() (err error) {
	// 配置文件
	viper.SetConfigFile("./config.yml") // 指定配置文件路径
	err = viper.ReadInConfig()          // 读取配置信息
	if err != nil {                     // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	return err
}
