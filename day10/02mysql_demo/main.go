package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

//查询单条记录
func queryOne(id int) {
	var u1 user
	sqlStr := `select id,name,age from user where id=?;`
	err := db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	if err != nil {
		return
	}
	fmt.Printf("u1:%#v\n", u1)
}

func queryMore(id int) {
	sqlStr := `select id,name,age from user where id > ?;`
	rowsObj, err := db.Query(sqlStr, id)
	if err != nil {
		return
	}
	defer rowsObj.Close()
	for rowsObj.Next() {
		var u1 user
		err := rowsObj.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			return
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

func insertUser(name string, age int) {
	sqlStr := `insert into user (name,age) values (?,?);`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId err:", err)
		return
	}
	fmt.Println("id:", id)

}

func updateRows(newAge, id int) {
	sqlStr := `update user set age = ? where id = ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Println("update err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err:", err)
		return
	}
	fmt.Printf("updated %d row(s)\n", n)
}

func deleteRow(id int) {
	sqlStr := `delete from user where id = ?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err:", err)
		return
	}
	fmt.Printf("delete %d row(s)\n", n)
}

func prepareInsert(m map[string]int) {
	sqlStr := `insert into user (name,age) values (?,?);`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare err:", err)
		return
	}
	defer stmt.Close()

	for k, v := range m {
		ret, err := stmt.Exec(k, v)
		if err != nil {
			fmt.Println("insert err:", err)
			return
		}
		id, err := ret.LastInsertId()
		if err != nil {
			fmt.Println("LastInsertId err:", err)
			return
		}
		fmt.Println("id:", id)
	}
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("Databases connect successful!")
	//queryOne(1)
	//queryMore(0)
	//insertUser("暗势",19)
	//updateRows(123,3)
	//deleteRow(5)

	//批量插入多条数据的时候：预编译
	var m = map[string]int{
		"惠子":   200,
		"色即是空": 300,
		"撒旦":   300,
		"阿松大":  300,
	}
	prepareInsert(m)
}
