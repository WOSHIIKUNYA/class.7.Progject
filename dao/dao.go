package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func Open() {
	dp, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	if err != nil {
		return
	}
	err = dp.Ping()
	if err != nil {
		return
	}
	database = dp
}
