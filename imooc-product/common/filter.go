// 拦截器服务
package common

import "net/http"

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
