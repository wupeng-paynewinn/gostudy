package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// mysql 语句：
/*	DDL :data defined 数据定义
	DML :data manipulation language 数据操纵
	DCL:data control language 数据控制
	TCL：transaction language 事务控制语言
*/

//存储引擎
/*
常见的存储引擎：MylSAM，InnoDB
*/
//func main() {
//	// DSN:Data Source Name
//	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()  // 注意这行代码要写在上面err判断的下面
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//}

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确

	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("Databases connect successful!")
}
