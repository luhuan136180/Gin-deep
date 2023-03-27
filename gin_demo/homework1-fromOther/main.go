package main

//
//import (
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//// 处理错误的方法
//
//// internal error
//type Er struct {
//	Code    int
//	Message string
//}
//
//func (e *Er) Error() string {
//	return e.Message
//}
//
//// api error
//type Err struct {
//	Code    int
//	Message string
//	Err     error
//}
//
//func (err *Err) Error() string {
//	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
//}
//
//func New(er *Er, err error) *Err {
//	return &Err{Code: er.Code, Message: er.Message, Err: err}
//}
//
//// main.go
//type userRegisterInfo struct {
//	Name          string `json:"name"`
//	Age           int    `json:"age"`
//	Passwd        string `json:"passwd"`
//	Gander        string `json:"gander"`
//	ConformPasswd string `json:"conform_passwd"`
//}
//
//type userLoginInfo struct {
//	Name   string `json:"name"`
//	Passwd string `json:"passwd"`
//}
//
//var db *sql.DB
//
//const dsn = "root:123456@tcp(127.0.0.1:3306)/sql_test"
//
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Hello A用户！")
//}
//
//func registerHandler(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	var answer = `{"status": "ok"}`
//	userInfo := &userRegisterInfo{}
//	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
//	//r.ParseForm()
//	//fmt.Println(r.PostForm)
//	//fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("passwd"), r.PostForm.Get("age"), r.PostForm.Get("gander"))
//	// 2. 请求类型是application/json时从r.Body读取数据
//	b, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Printf("read request.Body failed, err:%v\n", err)
//		return
//	}
//	fmt.Println(string(b))
//	err = json.Unmarshal(b, userInfo)
//	if err != nil {
//		fmt.Printf("解析参数时出错：%v\n", err.Error())
//		return
//	}
//	err = Register(userInfo)
//	if err != nil {
//		err = errnum.New(&errnum.Er{40001, "用户注册时出错"}, err)
//		answer = `{"status": "error"}`
//		w.Write([]byte(answer + err.Error()))
//		return
//	} else {
//		w.Write([]byte(answer))
//		return
//	}
//}
//
//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	var answer = `{"status": "ok", "data": "登录成功！"}`
//	userInfo := &userLoginInfo{}
//
//	b, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Printf("read request.Body failed, err:%v\n", err)
//		return
//	}
//	fmt.Println(string(b))
//	err = json.Unmarshal(b, userInfo)
//	if err != nil {
//		fmt.Printf("解析参数时出错：%v\n", err.Error())
//		return
//	}
//
//	err = Login(userInfo)
//	if err != nil {
//		err = errnum.New(&errnum.Er{40002, "用户登录时出错"}, err)
//		answer = `{"status": "error"}`
//		w.Write([]byte(answer + err.Error()))
//		return
//	} else {
//		w.Write([]byte(answer))
//		return
//	}
//
//	w.Write([]byte(answer))
//}
//
//func Register(userInfo *userRegisterInfo) (err error) {
//	// 校验用户输入的两次密码是否一致
//	if userInfo.Passwd != userInfo.ConformPasswd {
//		err = errnum.New(&errnum.Er{50001, "请保证两次输入密码的一致性"}, nil)
//		return err
//	}
//
//	// 连接数据库
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50002, "打开数据库连接失败！"}, err)
//		return err
//	}
//	// 测试连接数据库
//	err = db.Ping()
//	if err != nil {
//		err = errnum.New(&errnum.Er{50003, "连接数据库失败"}, err)
//		return err
//	}
//
//	// 查询当前用户是否已经存在
//	existsSqlStr := "select id from users where name = ?"
//	stmt, err := db.Prepare(existsSqlStr)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50004, "预处理查询SQL失败"}, err)
//		return err
//	}
//
//	defer stmt.Close()
//	rows, err := stmt.Query(userInfo.Name)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50005, "查询失败"}, err)
//		return err
//	}
//
//	if rows.Next() {
//		err = errnum.New(&errnum.Er{50006, "注册用户已存在"}, nil)
//		return err
//	}
//
//	// 如果注册用户不存在，则插入用户信息
//	insertSqlStr := "insert into users (name, passwd, age, gander) values (?, ?, ?, ?)"
//	stmt, err = db.Prepare(insertSqlStr)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50007, "预处理插入SQL失败"}, err)
//		return err
//	}
//	_, err = stmt.Exec(userInfo.Name, userInfo.Passwd, userInfo.Age, userInfo.Gander)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50008, "插入用户信息失败"}, err)
//		return err
//	}
//
//	return nil
//}
//
//func Login(userInfo *userLoginInfo) (err error) {
//	var u userLoginInfo
//	// 创建数据库连接
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50002, "打开数据库连接失败！"}, err)
//		return err
//	}
//
//	err = db.Ping()
//	if err != nil {
//		err = errnum.New(&errnum.Er{50003, "连接数据库失败"}, err)
//		return err
//	}
//
//	// 校验查询用户是否存在；如果存在，则校验用户的密码是否正确
//	searchSql := "select name, passwd from users where name = ?"
//
//	stmt, err := db.Prepare(searchSql)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50004, "预处理查询SQL失败"}, err)
//		return err
//	}
//	defer stmt.Close()
//	rows, err := stmt.Query(userInfo.Name)
//	if err != nil {
//		err = errnum.New(&errnum.Er{50005, "查询失败"}, err)
//		return err
//	}
//	if !rows.Next() {
//		err = errnum.New(&errnum.Er{51001, "登录用户不存在，请先注册"}, nil)
//		return err
//	}
//	err = rows.Scan(&u.Name, &u.Passwd)
//	if err != nil {
//		err = errnum.New(&errnum.Er{51002, "解析单行数据出错"}, err)
//		return err
//	}
//
//	if u.Passwd != userInfo.Passwd {
//		err = errnum.New(&errnum.Er{51003, "用户输入的密码有误"}, err)
//		return err
//	}
//
//	return nil
//}
//
//func main() {
//	http.HandleFunc("/index", indexHandler)
//	http.HandleFunc("/register", registerHandler)
//	http.HandleFunc("/login", loginHandler)
//	err := http.ListenAndServe(":9999", nil)
//	if err != nil {
//		fmt.Printf("http server failed, err:%v\n", err)
//		return
//	}
//}
