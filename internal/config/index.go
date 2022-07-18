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
	// 默认配置文件
	defaultIni := "./config/index.ini"

	// 解析当前环境配置文件
	var env = flag.String("env", "dev", "app env")
	flag.Parse()
	envIni := "./config/" + *env + ".ini"

	cfg, err := ini.LoadSources(ini.LoadOptions{
		// 跳过无法识别的数据行
		SkipUnrecognizableLines: true,
		// 出现没有具体值的布尔类型的键，这些键的值为 true
		AllowBooleanKeys: true,
	}, defaultIni, envIni)
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		return
	}
	config.AppEnv = *env
	err = cfg.MapTo(config)

	if err != nil {
		fmt.Println("映射配置文件信息失败: ", err)
		return
	}
}

func GetConfig() Config {
	return *config
}
