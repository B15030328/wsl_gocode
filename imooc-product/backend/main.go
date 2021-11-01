package main

import (
	"imooc-product/backend/web/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {

	app := iris.New()

	app.Logger().SetLevel("debug")

	// 注册模板
	tmplate := iris.HTML("../backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	// 设置模板目标
	app.HandleDir("/assets", "../backend/web/assets")
	// app.stat
	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 采用mvc模型
	// ctx, cancle := context.WithCancel(context.Background())
	// defer cancle()

	// 初始化product mvc路由
	productParty := app.Party("/product")
	product := mvc.New(productParty)

	// 初始化services并注册
	// productSerivce := services.NewProductService()
	// product.Register(ctx, productSerivce)

	product.Handle(new(controllers.ProductController))

	// 初始化order mvc路由
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Handle(new(controllers.OrderController))

	// 启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
