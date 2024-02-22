package initialize

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-backend/config"
	"go-backend/global"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.GlobConfig = serverConfig
	color.Blue("Config inits successfully\n")

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		color.Yellow("config file changed:", in.Name, "\n")
		// 重载配置
		if err := v.Unmarshal(&serverConfig); err != nil {
			fmt.Println(err)
		}
	})

}
