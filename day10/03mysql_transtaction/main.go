package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return nil
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("begin err:", err)
		return
	}
	fmt.Println("事务开启成功!")
	sqlStr1 := `update user set age=age-2 where id=1;`
	//sqlStr2 := `update user set age=age+2 where id=2;`
	//测试失败回滚
	sqlStr2 := `update xxx set age=age+2 where id=2;`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		fmt.Println("exec 1 err:", err)
		err = tx.Rollback()
		if err != nil {
			fmt.Println("rollback failed")
			return
		}
		return
	}
	fmt.Println("1 success")

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		fmt.Println("exec 2 err:", err)
		err = tx.Rollback()
		if err != nil {
			fmt.Println("rollback failed")
			return
		}
		return
	}
	fmt.Println("2 success")
	fmt.Println("事务执行成功!")

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit failed")
		err = tx.Rollback()
		if err != nil {
			fmt.Println("rollback failed")
			return
		}
		return
	}
	fmt.Println("事务提交成功!")

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("Databases connect successful!")
	transactionDemo()
}
