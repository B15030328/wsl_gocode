package mysql

import (
	"imooc-product/models/define"
	"log"

	"gorm.io/gorm"
)

type OrderInstance struct {
	db *gorm.DB
}

func NewOrderInstance() *OrderInstance {
	return &OrderInstance{
		db: defaultDb,
	}
}

// Insert(*define.Order) (int64,error)
// Delete(int64) bool
// Update(*define.Order) error
// SelectByKey (int64) (*define.Order,error)
// SelectAll ()([]*define.Order,error)
// SelectAllWithInfo()(map[int]map[string]string,error)

func (ins *OrderInstance) Insert(order *define.Order) (int64, error) {
	err := ins.db.Create(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Panic("order insert error,err:", err)
		return 0, err
	}
	return order.ID, nil
}

func (ins *OrderInstance) Delete(id int64) bool {
	err := ins.db.Delete(&define.Order{}, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Panic("order delete error,err:", err)
		return false
	}
	return true
}

func (ins *OrderInstance) Update(order *define.Order) error {
	err := ins.db.Save(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Panic("order update error,err:", err)
		return err
	}
	return nil
}

func (ins *OrderInstance) SelectByKey(id int64) (*define.Order, error) {
	var order *define.Order
	err := ins.db.First(&order, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Panic("order SelectByKey error,err:", err)
		return nil, err
	}
	return order, nil
}

func (ins *OrderInstance) SelectAll() ([]*define.Order, error) {
	var orders = make([]*define.Order, 0)
	err := ins.db.Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Panic("order SelectAll error,err:", err)
		return nil, err
	}
	return orders, nil
}

func (ins *OrderInstance) SelectAllWithInfo() ([]*define.Order, error) {
	var orders = make([]*define.Order, 0)
	res := ins.db.Find(&orders)
	err := res.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Panic("order insert error,err:", err)
		return nil, err
	}
	// data, _ := json.Marshal(orders)
	// fmt.Println(string(data))
	return orders, nil // todo
}
