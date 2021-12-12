package mysql

import (
	"imooc-product/models/define"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type Instance struct {
	DB *gorm.DB
}

func NewInstance() *Instance {
	db := defaultDb
	return &Instance{
		DB: db,
	}
}

// Insert(product *define.Product) (int64, error)
// Delete(id int64) bool
// Update(product *define.Product) error
// SelectByKey(productID int64) (productResult *define.Product, err error)
// SelectAll() (productArray []*define.Product, err error)

func (ins *Instance) Insert(product *define.Product) (int64, error) {
	result := ins.DB.Create(&product)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Panic("product insert error,err:", result.Error)
		return 0, result.Error
	}
	return product.ID, nil
}

func (ins *Instance) Delete(id int64) bool {
	result := ins.DB.Delete(&define.Product{}, id)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Panic("product delete error,err:", result.Error)
		return false
	}
	return true
}

func (ins *Instance) Update(product *define.Product) error {
	result := ins.DB.Save(&product)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Panic("product Update error,err:", result.Error)
		return result.Error
	}
	return nil
}

//根据商品ID查询商品
func (ins *Instance) SelectByKey(productID int64) (productResult *define.Product, err error) {
	result := ins.DB.First(&productResult, productID)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Panic("product SelectByKey error,err:", result.Error)
		return nil, result.Error
	}
	return productResult, nil
}

//获取所有商品
func (ins *Instance) SelectAll() (productArray []*define.Product, err error) {
	productArray = make([]*define.Product, 0)
	result := ins.DB.Find(&productArray)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Panic("product SelectAll error,err:", result.Error)
		return nil, result.Error
	}
	return productArray, nil
}

func (ins *Instance) SubProductNum(productID int64) error {
	result := ins.DB.Where("ID = ?", strconv.FormatInt(productID, 10)).Update("productNum", "productNum-1")

	if result.Error != nil {
		return result.Error
	}
	return result.Error
}
