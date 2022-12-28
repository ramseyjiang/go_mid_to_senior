package viperpkg

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const CfgName = "env"
const CfgType = "toml"
const CfgJoint = "."

type Config struct {
	AppName   string
	LogLevel  string
	MysqlConf MysqlConf
	RedisConf redisConf
}

type MysqlConf struct {
	ip       string
	user     string
	password string
	database string
	port     int64
}

type redisConf struct {
	ip   string
	port int64
}

// Entry run this file should go this folder, after that "go run config.go"
func Entry() (config *Config) {
	viper.SetConfigName(CfgName) // load config from which file name.
	viper.SetConfigType(CfgType) // load config from which file type
	viper.AddConfigPath(CfgJoint)
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	// if you declare uses "var config map[string]string", it will make "panic: assignment to entry in nil map"
	// so here you have to initialize the map using the make function (or a map literal) before you can add any elements.
	config = &Config{}
	// convert interface to string
	config.AppName = fmt.Sprintf("%v", viper.Get("app_name"))
	config.LogLevel = fmt.Sprintf("%v", viper.Get("log_level"))

	config.MysqlConf.ip = fmt.Sprintf("%v", viper.Get("mysql.ip"))
	config.MysqlConf.database = fmt.Sprintf("%v", viper.Get("mysql.database"))
	config.MysqlConf.user = fmt.Sprintf("%v", viper.Get("mysql.user"))
	config.MysqlConf.password = fmt.Sprintf("%v", viper.Get("mysql.password"))
	config.MysqlConf.port = (viper.Get("mysql.port")).(int64) // convert interface to int64

	config.RedisConf.ip = fmt.Sprintf("%v", viper.Get("redis.ip"))
	config.RedisConf.port = (viper.Get("redis.port")).(int64)

	return
}
