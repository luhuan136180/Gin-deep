package main

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type User struct {
	Id   int
	Age  int
	Name string
}

func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := DB.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In("INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users...,
	)
	fmt.Println(query)
	fmt.Println(args)
	_, err := DB.Exec(query, args...)
	return err
}

var DB *sqlx.DB

func queryRowDemo() {
	sqlStr := "select id,name,age from user where id=?"
	var u User
	err := DB.Get(&u, sqlStr, 2)
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id,name,age from user where id > ?"
	var users []User
	err := DB.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func insertRowDemo() {
	sqlStr := "insert into user(name,age) values (?,?)"
	ret, err := DB.Exec(sqlStr, "沙河", "19")
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d .\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age = ? where id =?"
	ret, err := DB.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed ,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func initDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/go_web program_01"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	return
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	fmt.Printf("INIT Db Success...\n")

	queryRowDemo()
	queryMultiRowDemo()
}
