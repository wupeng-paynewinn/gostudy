package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确

	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return nil
}

type user struct {
	ID   int
	Name string `json:"name" db:"name" ini:"name"`
	Age  int
}

func queryOne() {
	sqlStr := `select id,name,age from user where id = ?`
	var u1 user
	err := db.Get(&u1, sqlStr, 1)
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u1.ID, u1.Name, u1.Age)
}

func queryMulti() {
	sqlStr := `select id,name,age from user `
	var u []user
	err := db.Select(&u, sqlStr)
	if err != nil {
		fmt.Println("select err:", err)
		return
	}
	fmt.Printf("u:%#v\n", u)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("Databases connect successful!")
	queryMulti()
}
