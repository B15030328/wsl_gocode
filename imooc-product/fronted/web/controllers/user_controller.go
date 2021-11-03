package controllers

import (
	"imooc-product/encrypt"
	"imooc-product/models/define"
	"imooc-product/models/mysql"
	"imooc-product/services"
	"imooc-product/tool"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

// BeforeActivation 会被调用一次，在控制器适应主应用程序之前
// 并且当然也是在服务运行之前
func (u *UserController) BeforeActivation(b mvc.BeforeActivation) {
	u.Service.Ins = *mysql.NewUserInstance()
}

func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	//ozzo-validation
	user := &define.User{
		UserName:     userName,
		NickName:     nickName,
		HashPassword: password,
	}

	_, err := c.Service.AddUser(user)
	c.Ctx.Application().Logger().Debug(err)
	if err != nil {
		c.Ctx.Redirect("/user/error")
		return
	}
	c.Ctx.Redirect("/user/login")
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (c *UserController) PostLogin() mvc.Response {
	//1.获取用户提交的表单信息
	var (
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	//2、验证账号密码正确
	user, isOk := c.Service.IsPwdSuccess(userName, password)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	//3、写入用户ID到cookie中
	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	// 用cookie代替session 不需要再下面的写法
	// session := c.Session.Start(c.Ctx)
	// session.Set("userID", strconv.FormatInt(user.ID, 10))
	uidByte := []byte(strconv.FormatInt(user.ID, 10))
	uidString, err := encrypt.EnPwdCode(uidByte)
	if err != nil {
		log.Panic(err)
	}
	// 写入用户浏览器
	tool.GlobalCookie(c.Ctx, "sign", uidString)
	return mvc.Response{
		Path: "/product/",
	}

}
