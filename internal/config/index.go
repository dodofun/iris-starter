package config

import (
	"flag"
	"fmt"

	"gopkg.in/ini.v1"
)

var config = &Config{}

// 配置信息
type Config struct {
	AppName string
	AppEnv  string
	Server
	Log
	Mysql
	Redis
}

func init() {
	// 解析当前环境
	var env = flag.String("env", "dev", "app env")
	flag.Parse()
	envIni := "./config/" + *env + ".ini"

	cfg, err := ini.Load("./config/index.ini", envIni)
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		return
	}
	config.AppEnv = *env
	err = cfg.MapTo(config)
	fmt.Println("AppEnv: " + config.AppEnv)
	fmt.Println("HttpPort:" + config.HttpProtocol)
	fmt.Println("Mysql:", config.Mysql.Read.Address, config.Redis.Write.Password)
	if err != nil {
		fmt.Println("映射配置文件信息失败: ", err)
		return
	}
}

func GetConfig() Config {
	return *config
}
