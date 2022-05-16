package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//连接数据库
func initdb() *gorm.DB {
	conn, err := gorm.Open(
		mysql.Open("jiuxia:!zzh020502@tcp(rm-bp15zhrxyp3qcn7bfto.mysql.rds.aliyuncs.com:3306)/douyin"))
	if err != nil {
		panic(err)
	}
	return conn
}

//初始化连接和线程池的设置 最大打开:30  最大闲置:10
func Setup() {
	db = initdb()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(30) // 设置最大打开的连接数
	sqlDB.SetMaxIdleConns(10) // 设置最大闲置的连接数
}

//获取查询实例
func GetDB() *gorm.DB {
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		Setup()
	}
	return db
}
