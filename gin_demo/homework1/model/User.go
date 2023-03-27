package model

import "gin_demo/homework1/dao"

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Money    float64 `json:"money"`
	Age      int     `json:"age"`
}

//注册
func InsertUser(u *User) (err error) {
	err = dao.DB.Create(u).Error
	return
}

//通过id查询
func SelectUser(id int) (u *User, err error) {
	u = new(User) //在输出部分创立的变量在方法体重需要赋值
	err = dao.DB.Where("id=?", id).Find(u).Error
	return
}

//通过name查询
func FindName(name string) (u *User, err error) {
	u = new(User) //在输出部分创立的变量在方法体重需要赋值
	//var data User
	err = dao.DB.Model(User{}).Where("name=?", name).First(u).Error
	return
}

//更新
func UpdateUser(u *User) (err error) {
	err = dao.DB.Save(u).Error
	return
}
func DeleteUser(id int) (err error) {
	err = dao.DB.Where("id=?", id).Delete(User{}).Error
	return
}
