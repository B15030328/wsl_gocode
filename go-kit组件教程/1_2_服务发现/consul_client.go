package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

/*
服务发现
*/

func main() {
	// var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500" //consul server

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return
	}
	services, _, err := client.Health().Service("userservice", "primary", true, nil)
	if err != nil {
		logrus.Warn("error retrieving instances from Consul: %v", err)
	}
	// lastIndex = metainfo.LastIndex
	addrs := map[string]struct{}{}
	for _, service := range services {
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}
	fmt.Println(addrs)
}
