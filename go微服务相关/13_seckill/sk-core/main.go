package main

import (
	"fmt"
	"seckill/pkg/bootstrap"
)

/*
sk-app
实现秒杀业务系统
*/

func main() {

	fmt.Println(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)

}
