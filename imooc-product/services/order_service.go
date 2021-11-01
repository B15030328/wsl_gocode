package services

import (
	"imooc-product/models/define"
	"imooc-product/models/mysql"
)

type OrderService struct {
	Ins mysql.OrderInstance
}

//初始化函数
func NewOrderService() *OrderService {
	ins := mysql.NewOrderInstance()
	return &OrderService{Ins: *ins}
}

func (o *OrderService) GetProductByID(productID int64) (*define.Order, error) {
	return o.Ins.SelectByKey(productID)
}

func (o *OrderService) DeleteOrderByID(orderID int64) (isOk bool) {
	isOk = o.Ins.Delete(orderID)
	return
}

func (o *OrderService) UpdateOrder(order *define.Order) error {
	return o.Ins.Update(order)
}

func (o *OrderService) InsertOrder(order *define.Order) (orderID int64, err error) {
	return o.Ins.Insert(order)
}

func (o *OrderService) GetAllOrder() ([]*define.Order, error) {
	return o.Ins.SelectAll()
}

func (o *OrderService) GetAllOrderInfo() ([]*define.Order, error) {
	return o.Ins.SelectAllWithInfo()
}
