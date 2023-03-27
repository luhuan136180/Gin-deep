package main

import (
	"fmt"
	"gin_demo/homework1/dao"
	"gin_demo/homework1/model"
	"gin_demo/homework1/routers"
)

/*	结合net/http和database/sql实现一个使用MySQL存储用户信息的注册及登陆的简易web程序。
建立数据库
建立一张用户信息表（用户名，密码，金额，年龄，）
*/
//第一步完成表的信息录入，修改，删除(完成)
//第二步：建立路由器，分别完成信息录入，修改，更新，删除工作
//建立状态表

func main() {
	err := dao.InitMysql()
	if err != nil {
		fmt.Println("1")
		panic(err)
		return
	}
	//模型绑定
	dao.DB.AutoMigrate(&model.User{})

	//启动引擎
	r := routers.InitRouter()
	r.Run()

	////查找单个
	//u2, err := model.SelectUser(3)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Println(u2)
	////修改
	//u2.Money = 150.0
	//model.UpdateUser(u2)
	//u3, err := model.SelectUser(3)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Println(u3)
	////删除
	//model.DeleteUser(4)
	//u, err := model.FindName("mhx")
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Println(u)
}
