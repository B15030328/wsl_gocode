package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var defaultDb = InitDB()

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/seckillmall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("连接数据库出错，err: ", err)
	}
	// productArray := make([]*define.Product, 0)
	// fmt.Println("111")
	// fmt.Println(productArray)
	// _ = db.Find(&productArray)
	// fmt.Println(productArray)
	return db
}
