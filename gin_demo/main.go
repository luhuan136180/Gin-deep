package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Age  int
	Name string
}

var (
	Db *sql.DB
)

func initDb() (err error) {
	dsn := "root:root@tcp(localhost:3306)/go_web program_01"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func queryMultiRowDemo(id int) {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := Db.Query(sqlStr, id)
	// 查询多条数据示例
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}

}
func queryRowDemo(id int) {
	var u User
	sqlStr := "select id, name, age from user where id=?"
	err := Db.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

func insertRowDemo(u *User) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := Db.Exec(sqlStr, u.Name, u.Age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theId, err := ret.LastInsertId() //新插入的数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theId)
}
func UpdateRowDemo() {
	sqlStr := "update user set age=? where id=?"
	ret, err := Db.Exec(sqlStr, 39, 3) //按照？的先后顺序输入值
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func main() {
	err := initDb() //调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	defer Db.Close()
	queryRowDemo(1)
	queryMultiRowDemo(1)
	u := User{Age: 21, Name: "毛浩昕"}
	insertRowDemo(&u)
}
