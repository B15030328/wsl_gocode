package controllers

import (
	"encoding/json"
	"imooc-product/models/define"
	"imooc-product/models/mysql"
	"imooc-product/rabbitmq"
	"imooc-product/services"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.ProductService
	OrderService   services.OrderService
	Session        *sessions.Sessions
	RabbitMQ       *rabbitmq.RabbitMQ
}

var (
	// 静态资源生成的html保存目录
	htmlOutPath = "../fronted/web/htmlProductShow/"
	// 静态文件模板目录
	templatePath = "../fronted/web/views/template/"
)

// BeforeActivation 会被调用一次，在控制器适应主应用程序之前
// 并且当然也是在服务运行之前
func (c *ProductController) BeforeActivation(b mvc.BeforeActivation) {
	c.ProductService.Ins = *mysql.NewInstance()
	c.OrderService.Ins = *mysql.NewOrderInstance()
	sess := sessions.New(sessions.Config{
		Cookie:  "AdminCookie",
		Expires: 600 * time.Minute,
	})
	c.Session = sess
	c.RabbitMQ = rabbitmq.NewRabbitMQSimple("prod")
}

func (p *ProductController) GetGenerateHtml() {
	productString := p.Ctx.URLParam("productID")
	productID, err := strconv.Atoi(productString)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	//1.获取模版
	contenstTmp, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//2.获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")

	//3.获取模版渲染数据
	product, err := p.ProductService.GetProductByID(int64(productID))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//4.生成静态文件
	generateStaticHtml(p.Ctx, contenstTmp, fileName, product)
}

//生成html静态文件
func generateStaticHtml(ctx iris.Context, template *template.Template, fileName string, product *define.Product) {
	//1.判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			ctx.Application().Logger().Error(err)
		}
	}
	//2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
	defer file.Close()
	template.Execute(file, &product)
}

//判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProductByID(4)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetOrder() []byte {
	productString := p.Ctx.URLParam("productID")
	userString := p.Ctx.GetCookie("uid")
	productID, err := strconv.ParseInt(productString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	// 采用消息队列的方式查询
	userID, err := strconv.ParseInt(userString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	message := define.NewMessage(userID, int64(productID))
	byteMessage, err := json.Marshal(message)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	// 生产者sample模式
	err = p.RabbitMQ.PublishSimple(string(byteMessage))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	// product, err := p.ProductService.GetProductByID(int64(productID))
	// if err != nil {
	// 	p.Ctx.Application().Logger().Debug(err)
	// }
	// var orderID int64
	// showMessage := "抢购失败！"
	// //判断商品数量是否满足需求
	// if product.ProductNum > 0 {
	// 	//扣除商品数量
	// 	product.ProductNum -= 1
	// 	err := p.ProductService.UpdateProduct(product)
	// 	if err != nil {
	// 		p.Ctx.Application().Logger().Debug(err)
	// 	}
	// 	//创建订单
	// 	userID, err := strconv.Atoi(userString)
	// 	if err != nil {
	// 		p.Ctx.Application().Logger().Debug(err)
	// 	}

	// 	order := &define.Order{
	// 		UserId:      int64(userID),
	// 		ProductId:   int64(productID),
	// 		OrderStatus: define.OrderSuccess,
	// 	}
	// 	//新建订单
	// 	orderID, err = p.OrderService.InsertOrder(order)
	// 	if err != nil {
	// 		p.Ctx.Application().Logger().Debug(err)
	// 	} else {
	// 		showMessage = "抢购成功！"
	// 	}
	// }
	return []byte("true")
	// return mvc.View{
	// 	Layout: "shared/productLayout.html",
	// 	Name:   "product/result.html",
	// 	Data: iris.Map{
	// 		"orderID":     orderID,
	// 		"showMessage": showMessage,
	// 	},
	// }

}
