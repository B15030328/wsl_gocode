package main

import (
	"imooc-product/fronted/middleware"
	"imooc-product/fronted/web/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	tmplate := iris.HTML("../fronted/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板
	app.HandleDir("/public", "../fronted/web/public")
	//访问生成好的html静态文件
	app.HandleDir("/html", "../fronted/web/htmlProductShow")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	userPro := mvc.New(app.Party("/user"))
	userPro.Handle(new(controllers.UserController))
	proParty := app.Party("/product")
	proParty.Use(middleware.AuthConProduct)
	pro := mvc.New(proParty)
	pro.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("localhost:8082"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
