package viperpkg

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const CfgName = "env"
const CfgType = "toml"
const CfgJoint = "."

// Entry run this file should go this folder, after that "go run config.go"
func Entry() {
	viper.SetConfigName(CfgName) // load config from which file name.
	viper.SetConfigType(CfgType) // load config from which file type
	viper.AddConfigPath(CfgJoint)
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))
}
