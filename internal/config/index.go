package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var config = &Config{}

// 配置信息
type Config struct {
	AppName string
	AppMode string
	Server
	Log
	Mysql
	Redis
}

func init() {
	fmt.Println("读取配置文件")
	cfg, err := ini.Load("./config/index.ini")
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		return
	}
	err = cfg.MapTo(config)
	fmt.Println("AppName: " + config.AppName)
	fmt.Println("HttpPort:" + config.HttpPort)
	fmt.Println("Mysql:", config.Mysql.Read.Address, config.Redis.Write.Password)
	if err != nil {
		fmt.Println("映射配置文件信息失败: ", err)
		return
	}
}

func GetConfig() Config {
	return *config
}
