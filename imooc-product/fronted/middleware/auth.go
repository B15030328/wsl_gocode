package middleware

import (
	"fmt"
	"imooc-product/encrypt"
	"log"

	"github.com/kataras/iris/v12"
)

func AuthConProduct(ctx iris.Context) {

	sign := ctx.GetCookie("sign")
	fmt.Println(sign)
	sign_decode, err := encrypt.DePwdCode(sign)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(sign_decode))
	if string(sign_decode) == "" {
		ctx.Application().Logger().Debug("必须先登录!")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug("已经登陆")
	ctx.Next()
}
