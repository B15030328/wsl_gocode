package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type User struct {
	Name string
	Sex  string
	Age  int
}

var Resume User

func init() {
	viper.AutomaticEnv()
	initDefault()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	if err := sub("User1111", &Resume); err != nil {
		log.Fatal("Fail to parse config", err)
	}
}

func initDefault() {
	//设置读取的配置文件
	viper.SetConfigName("resume_config")
	//添加读取的配置文件路径
	viper.AddConfigPath("config/")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	viper.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	viper.SetConfigType("yaml")
}

func sub(key string, value interface{}) error {
	log.Printf("配置文件的前缀为：%v", key)
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	return sub.Unmarshal(value)
}

func main() {
	fmt.Printf(
		"姓名: %s\n性别: %s \n年龄: %d \n",
		Resume.Name,
		Resume.Sex,
		Resume.Age,
	)
}
