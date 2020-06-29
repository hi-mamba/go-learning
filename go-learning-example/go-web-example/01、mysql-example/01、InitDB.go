package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 主意这里
)

const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "example"
)

//Db数据库连接池
var SqlDB *sql.DB

//注意方法名大写，就是public

func InitDB() SqlDB.DB {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	//DB, _ = sql.Open("mysql", path)
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/example")
	//DB, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/example")
	fmt.Println("----", err)
	//设置数据库最大连接数
	SqlDB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	SqlDB.SetMaxIdleConns(10)
	//验证连接
	if err := SqlDB.Ping(); err != nil {
		fmt.Println("opon database fail", err)
		return sql.DB{}
	}
	fmt.Println("connect success")
	//defer SqlDB.Close()
	return sql.DB{};
}
