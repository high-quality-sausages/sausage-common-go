package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

const (
	_dsn = "%s:%s@tcp(%s)/%s?charset=utf8"
)

// NewMysql create a mysql object
func NewMysql(conf *mysql.Config) *sql.DB {
	path := fmt.Sprintf(_dsn, conf.User, conf.Passwd, conf.Addr, conf.DBName)
	db, _ := sql.Open("mysql", path)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		fmt.Printf("failed to open mysql, err:%v\n", err)
		return nil
	}
	return db
}
