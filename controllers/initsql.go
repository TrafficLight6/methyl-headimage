package controllers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "headimg:headimg@tcp(127.0.0.1:3306)/headimg")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db = database
}
