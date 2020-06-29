package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 主意这里
)

var DB *sql.DB


func InsertUser() bool {
	DB,_ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/example")
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`name`, `age`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail", err)
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec("张三", 33)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true

}

func SelectUser() {
	//执行查询语句
	rows, err := DB.Query("SELECT id,name from user ")
	if err != nil{
		fmt.Println("查询出错了")
	}

	//循环读取结果
	for rows.Next(){
		//将每一行的结果都赋值到一个user对象中
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Printf("rows id = %d, value = %s", id, name)

	}
	//return users
}

func main()  {
	InsertUser()
	SelectUser()
}
