// 拦截器服务
package common

import (
	"errors"
	"imooc-product/encrypt"
	"log"
	"net/http"
)

// 声明过滤处理函数
type FilterHandle func(rw http.ResponseWriter, r *http.Request) error

type Filter struct {
	filterMap map[string]FilterHandle
}

// 初始化filter函数
func NewFilter() *Filter {
	return &Filter{filterMap: make(map[string]FilterHandle)}
}

// 注册针对某个url的拦截器
func (f *Filter) RegisterFilterUrl(url string, handler FilterHandle) {
	f.filterMap[url] = handler
}

//根据Uri获取对应的handle
func (f *Filter) GetFilterHandle(uri string) FilterHandle {
	return f.filterMap[uri]
}

// 声明handle
type webHandle func(rw http.ResponseWriter, r *http.Request)

// 执行拦截器
func (f *Filter) Handle(webHandle webHandle) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		for url, handle := range f.filterMap {
			// 如果有对应的uri在拦截器中，则执行
			if url == r.RequestURI {
				err := handle(rw, r)
				if err != nil {
					rw.Write([]byte(err.Error()))
					return
				}
				break
			}
		}
		webHandle(rw, r)
	}
}

// 拦截器执行函数
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
