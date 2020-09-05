package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type config struct {
	Server Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	Username string
	Password string
	Host string
	Dbname string
}

var Conf *config

func init() {
	Conf = loadConfig()
}

func loadConfig() *config {
	dir, err := os.Getwd()
	if err != nil {
		panic("读取配置文件目录错误：" + err.Error())
	}
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("application")
	v.AddConfigPath(dir + "/config")
	v.SetDefault("server.port", "8080")
	v.SetDefault("database.host", "localhost:3306")
	err = v.ReadInConfig()
	if err != nil {
		panic("读取配置文件错误：" + err.Error())
	}
	var c config
	err = v.Unmarshal(&c)
	if err != nil {
		panic("解析配置文件错误：" + err.Error())
	}
	log.Printf("config : %v", c)
	return &c
}