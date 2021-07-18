package utils

import (
	"fmt"
	"github.com/HSczy/CLBackendGO/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper 使用Viper读取配置文件
func Viper() *viper.Viper {
	v := viper.New()
	//设置配置文件
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置文件错误，错误信息为:%v \n", err)
	}
	// 监听文件变化
	v.WatchConfig()
	//文件变化重新读取配置
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("检测到配置文件修改，修改内容为:%v \n", in)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Printf("解析配置文件错误，错误信息为：%s \n", err)
		}
	})
	// 读取配置文件
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Printf("解析配置文件错误，错误信息为：%s \n", err)
	}
	return v
}
