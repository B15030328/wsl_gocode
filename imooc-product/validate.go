// 分布式权限验证
package main

import (
	"fmt"
	"imooc-product/common"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// 用来存放用户信息
type AccessControl struct {
	sourceArray sync.Map
}

var (
	hostArray      = []string{"127.0.0.1", "127.0.0.1"}
	localHost      = "127.0.0.1"
	port           = "8081"
	hashConsistent *common.Consistent
	map1           sync.Map
	accessControl  = &AccessControl{sourceArray: map1}
)

// 获取对应的信息
func (m *AccessControl) GetNewRecord(uid int) interface{} {
	data, ok := m.sourceArray.Load(uid)
	if !ok {
		log.Fatal("get record err")
	}
	return data
}

// 新增记录
func (m *AccessControl) SetNewRecord(uid int) {
	m.sourceArray.Store(uid, "hello chory")
}

// 分布式权限验证：在用户数据应当所在节点上查找用户数据，找到则验证成功
func (m *AccessControl) GetDistributedRight(req *http.Request) bool {
	//获取用户UID
	uidCookie, err := req.Cookie("uid")
	if err != nil {
		log.Println("getDistributedRight get cookie err")
		return false
	}
	//采用一致性hash算法，根据用户ID，判断获取具体机器
	hostRequest, err := hashConsistent.Get(uidCookie.Value)
	if err != nil {
		log.Println("getDistributedRight get hostRequest err")
		return false
	}
	//判断是否为本机
	if hostRequest == localHost {
		return m.GetDataFromMap(uidCookie.Value)
	}
	return GetDataFromOtherMap(hostRequest, req)
}

//获取本机map，并且处理业务逻辑，返回的结果类型为bool类型
func (m *AccessControl) GetDataFromMap(uid string) (isOk bool) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return false
	}
	data := m.GetNewRecord(uidInt)

	//执行逻辑判断
	if data != nil {
		return true
	}
	return
}

//获取其它节点处理结果
func GetDataFromOtherMap(host string, request *http.Request) bool {
	//获取Uid
	uidPre, err := request.Cookie("uid")
	if err != nil {
		return false
	}
	//获取sign
	uidSign, err := request.Cookie("sign")
	if err != nil {
		return false
	}
	// 模拟接口访问
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+host+":"+port+"/check", nil)
	if err != nil {
		log.Println("模拟请求用户数据所在节点请求失败")
		return false
	}
	//手动指定，排查多余cookies
	cookieUid := &http.Cookie{Name: "uid", Value: uidPre.Value, Path: "/"}
	cookieSign := &http.Cookie{Name: "sign", Value: uidSign.Value, Path: "/"}
	req.AddCookie(cookieUid)
	req.AddCookie(cookieSign)
	// 获取返回结果
	resp, err := client.Do(req)
	if err != nil {
		log.Println("模拟请求用户数据所在节点返回失败")
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("模拟请求用户数据所在节点返回body失败")
		return false
	}
	if resp.StatusCode == 200 && string(body) == "true" {
		return true
	}
	return false

}

func main() {
	// 负载均衡器设置
	// 采用hash一致性算法
	hashConsistent := common.NewConsistent()
	// 添加节点
	for _, v := range hostArray {
		hashConsistent.Add(v)
	}

	// 注册过滤器
	filter := common.NewFilter()
	// 注册拦截器
	filter.RegisterFilterUrl("/check", common.Auth)
	log.Println("service start success")
	// 启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	http.ListenAndServe(":8083", nil)
}

// 执行正常业务逻辑
func Check(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("执行check")
}
