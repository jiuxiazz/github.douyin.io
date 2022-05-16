package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"sync"
)

type User struct {
	ID   int
	Name string
}

// 线程池的测试
func main() {
	dao.Setup()
	//设置两个线程
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		db := dao.GetDB()
		var users []User
		//查询users表中的id字段
		db.Select("id").Find(&users)
		//打印
		fmt.Println(users)
		// 关闭该线程
		defer wg.Done()
	}()
	go func() {
		db := dao.GetDB()
		var users []User
		db.Select("name").Find(&users)
		fmt.Println(users)
		defer wg.Done()
	}()
	// 等待计数器归0
	wg.Wait()
}
