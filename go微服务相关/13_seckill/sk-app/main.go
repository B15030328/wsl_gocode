package main

import (
	"fmt"
	"seckill/pkg/bootstrap"
	"seckill/sk-app/setup"
)

/*
sk-app
实现秒杀业务系统
*/

func main() {
	// 初始化zookeeper 加载秒杀数据
	setup.InitZk()

	// 初始化redis 获取黑名单 启动redis协程
	setup.InitRedis()

	// 初始化服务
	// setup.InitServer()
	fmt.Println(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)

}
