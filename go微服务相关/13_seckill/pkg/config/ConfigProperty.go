package conf

import (
	"sync"

	"github.com/samuel/go-zookeeper/zk"
)

var (
	SecKill SecKillConf
	Zk      ZookeeperConf
)

type SecKillConf struct {
	RWBlackLock       sync.RWMutex
	SecProductInfoMap map[int]*SecProductInfoConf
}

type ZookeeperConf struct {
	ZkConn        *zk.Conn
	SecProductKey string // 商品键
}

// 商品信息配置
type SecProductInfoConf struct {
	ProductId         int     `json:"product_id"`           //商品ID
	StartTime         int64   `json:"start_time"`           //开始时间
	EndTime           int64   `json:"end_time"`             //结束时间
	Status            int     `json:"status"`               //状态
	Total             int     `json:"total"`                //商品总数量
	Left              int     `json:"left"`                 //商品剩余数量
	OnePersonBuyLimit int     `json:"one_person_buy_limit"` //单个用户购买数量限制
	BuyRate           float64 `json:"buy_rate"`             //购买频率限制
	SoldMaxLimit      int     `json:"sold_max_limit"`
	// todo: error
	SecLimit *SecLimit `json:"sec_limit"` //限速控制
}

//每秒限制
type SecLimit struct {
	count   int   //次数
	preTime int64 //上一次记录时间
}

//当前秒的访问次数
func (p *SecLimit) Count(nowTime int64) (curCount int) {
	if p.preTime != nowTime {
		p.count = 1
		p.preTime = nowTime
		curCount = p.count
		return
	}

	p.count++
	curCount = p.count
	return
}

func (p *SecLimit) Check(nowTime int64) int {
	if p.preTime != nowTime {
		return 0
	}
	return p.count
}
