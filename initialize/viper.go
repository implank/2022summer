package initialize

import (
	"2022summer/global"
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() {
	global.VP = viper.New()
	global.VP.SetConfigFile("./config.yml") // 指定配置文件路径
	err := global.VP.ReadInConfig()         // 读取配置信息
	if err != nil {                         // 读取配置信息失败
		panic(fmt.Errorf("读取配置文件失败, viper 出问题啦: %s \n", err))
		return
	}
	global.VP.WatchConfig()
	return
}
