package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConf)

type AppConf struct {
	*App `mapstructure:"app"`
}
type App struct {
	Port     string `mapstructure:"port"`
	Env      string `mapstructure:"env"`
	App_name string `mapstructure:"app_name"`
	Version  int    `mapstructure:"version"`
}

func Init() (err error) {
	viper.SetConfigFile("../config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("vip init err :", err)
		return
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return

}
