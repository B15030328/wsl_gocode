package setup

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	conf "seckill/pkg/config"

	"github.com/samuel/go-zookeeper/zk"
)

//初始化Zookeeper
func InitZk() {
	var hosts = []string{"localhost:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println("zookeeper connection error", err)
		return
	}
	fmt.Println("连接成功")

	conf.Zk.ZkConn = conn
	conf.Zk.SecProductKey = "/product"
	// 加载秒杀商品信息
	loadSecConf(conn)
}

// 加载秒杀商品信息
func loadSecConf(conn *zk.Conn) {
	log.Printf("Connect zk success %s", conf.Zk.SecProductKey)
	v, _, err := conn.Get(conf.Zk.SecProductKey)
	if err != nil {
		log.Println("get product info failed", err)
		return
	}
	log.Println("get product success")

	var secProductInfoConf []*conf.SecProductInfoConf
	err1 := json.Unmarshal(v, &secProductInfoConf)
	if err1 != nil {
		log.Printf("Unmsharl second product info failed, err : %v", err1)
	}

	// 更新秒杀商品信息
	updateSecProductInfo(secProductInfoConf)
}

// 更新秒杀商品信息
func updateSecProductInfo(secProductInfo []*conf.SecProductInfoConf) {
	var tmp = make(map[int]*conf.SecProductInfoConf, 1024)
	for _, v := range secProductInfo {
		log.Printf("updateSecProductInfo %v", v)
		tmp[v.ProductId] = v
	}
	conf.SecKill.RWBlackLock.Lock()
	conf.SecKill.SecProductInfoMap = tmp
	conf.SecKill.RWBlackLock.Unlock()
}
