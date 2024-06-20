package main

import (
	"chatim/conf"
	"chatim/dao/sqlites"
	"chatim/loger"
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
	// if err := mysqld.Init(); err != nil {
	// 	fmt.Println("mysql init err ", err)
	// 	return
	// }
	// defer mysqld.Close()
	//初始化postgres数据库
	// 	if err := postgres.Init(); err != nil {
	// 		fmt.Println("postgres init err ", err)
	// 		return
	// 	}
	// defer	postgres.Close()
	// fmt.Println("dds")
	if err := sqlites.Init(); err != nil {
		fmt.Println("init db err:", err)
	}
	defer sqlites.Close()
	sqlites.Dksd()
}
