// 分布式验证
package main

import (
	"errors"
	"fmt"
	"imooc-product/common"
	"imooc-product/encrypt"
	"log"
	"net/http"
)

//执行正常业务逻辑
func Check(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("执行check")
}

//统一验证拦截器，每个接口都需要提前验证
func Auth(rw http.ResponseWriter, r *http.Request) error {
	log.Println("执行验证")
	err := checkUserInfo(r)
	if err != nil {
		return err
	}
	return nil
}

// cookie验证
func checkUserInfo(r *http.Request) error {
	// 获取cookie中的uid
	uidCookie, err := r.Cookie("uid")
	if err != nil {
		log.Println("get uid cookie error")
		return errors.New("get uid cookie error")
	}
	// 获取cookie中的sign
	signCookie, err := r.Cookie("sign")
	if err != nil {
		log.Println("get sign cookie error")
		return errors.New("get sign cookie error")
	}
	sign, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		log.Println("decoder sign cookie error")
		return errors.New("decoder sign cookie error")
	}
	if checkInfo(uidCookie.Value, string(sign)) {
		log.Println("身份验证成功，uid", uidCookie.Value)
		return nil
	}
	return errors.New("身份校验失败！")
}

// 验证解密后的sign
func checkInfo(checkStr, signStr string) bool {
	return checkStr == signStr
}

func main() {
	// 1、注册过滤器
	filter := common.NewFilter()
	filter.RegisterFilterUrl("/check", Auth)
	log.Println("service start success")
	http.HandleFunc("/check", filter.Handle(Check))
	http.ListenAndServe(":8083", nil)
}
