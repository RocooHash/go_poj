package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Init() error {
	var err error

	DB, err = sql.Open("mysql", "root:123456@tcp(42.193.101.112:3306)/test?charset=utf8")
	if err != nil {
		return err
	}
	DB.SetMaxIdleConns(50)

	err = DB.Ping()
	if err != nil {
		return err
	} else {
		log.Println("Mysql Startup Normal!")
	}
	return nil
}
