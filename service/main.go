package main

import (
	"ChatRepository/service/conf"
	mysql "ChatRepository/service/dao"
	"ChatRepository/service/loger"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	// 1：加载配置
	if err := conf.Init(); err != nil {
		fmt.Println("viper init err ", err)
		return
	}

	// 2：初始化日志
	if err := loger.Init(); err != nil {
		fmt.Println("viper init err ", err)
		return
	}
	zap.L().Debug("log init success")
	// 初始化msyql数据库
	if err := mysql.Init(); err != nil {
		fmt.Println("mysql init err ", err)
		return
	}
	mysql.Close()
}
