package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

// 将变量注册到结构体，方便引入到html对应位置
type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

// 首页
func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// 检查是否存在cookie
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		// 模板中的变量
		p := &HomePage{Name: "chory"}

		// 模板解析
		t, err := template.ParseFiles("../template/home.html")
		if err != nil {
			log.Printf("Parsing template home.html error: %s", err)
			return
		}
		t.Execute(w, p)
		return
	}

	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/username", http.StatusFound)
		return
	}

}

// 用户页
func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		log.Println(err1, err2)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")
	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}
	t, e := template.ParseFiles("../template/userhome.html")
	if e != nil {
		log.Printf("Parsing userhome.html error: %s", e)
		return
	}
	t.Execute(w, p)
}

// api的方式透传
func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}
	request(apibody, w, r)
	defer r.Body.Close()
}

// proxy的方式透传
func proxyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, err := url.Parse("http://localhost:9000/")
	if err != nil {
		log.Println("prase url err,", err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
