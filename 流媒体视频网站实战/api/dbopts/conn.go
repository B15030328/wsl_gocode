package dbopts

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func Init() {
	log.Println("connect mysql...")
	dbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/streamvideo_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	// defer dbConn.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = dbConn.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	log.Println("connect mysql success")
}
