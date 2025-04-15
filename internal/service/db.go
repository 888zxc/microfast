package service

import (
	"database/sql"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
)

// 获取数据库连接池
func NewDBPool(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	// 设置连接池参数
	db.SetMaxOpenConns(100)            // 最大连接数
	db.SetMaxIdleConns(20)             // 最大空闲连接数
	db.SetConnMaxLifetime(time.Hour)   // 连接可重用的最大时间
	db.SetConnMaxIdleTime(30 * time.Minute) // 空闲连接保留时间
	
	return db, nil
}
