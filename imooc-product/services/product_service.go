package services

import (
	"imooc-product/models/define"
	"imooc-product/models/mysql"
)

type IProductService interface {
	GetProductByID(int64) (*define.Product, error)
	GetAllProduct() ([]*define.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *define.Product) (int64, error)
	UpdateProduct(product *define.Product) error
	SubNumberOne(int64) error
}

type ProductService struct {
	Ins mysql.Instance
}

//初始化函数
func NewProductService() IProductService {
	ins := mysql.NewInstance()
	return &ProductService{Ins: *ins}
}

func (p *ProductService) GetProductByID(productID int64) (*define.Product, error) {
	return p.Ins.SelectByKey(productID)
}

func (p *ProductService) GetAllProduct() ([]*define.Product, error) {
	return p.Ins.SelectAll()
}

func (p *ProductService) DeleteProductByID(productID int64) bool {
	return p.Ins.Delete(productID)
}

func (p *ProductService) InsertProduct(product *define.Product) (int64, error) {
	return p.Ins.Insert(product)
}

func (p *ProductService) UpdateProduct(product *define.Product) error {
	return p.Ins.Update(product)
}

func (p *ProductService) SubNumberOne(productID int64) error {
	return p.Ins.SubProductNum(productID)
}
