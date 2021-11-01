package models

import "imooc-product/models/mysql"

var product mysql.Instance

func Init() {
	product = *mysql.NewInstance()
}
