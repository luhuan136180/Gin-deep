package controller

import (
	"fmt"
	"gin_demo/homework1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册
func RegisterHandler(c *gin.Context) {
	//第一步从请求中获取注册用的表单数据
	//创建一个user实例接收数据
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

	} else {
		fmt.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"UserInfo": user,
		})
	}
	//检查用户表中是否已经存在该用户
	u, _ := model.FindName(user.Name)

	if u.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"Code": 1,
			"Msg":  "用户名已存在",
		})
		return
	}
	//若不存在，通过方法将实例数据数据库中存储

	err = model.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": 1,
			"Msg":  "注册失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code": 0,
		"Msg":  "注册成功",
	})
}

//登录
func LoginHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user, _ := model.FindName(name)
	fmt.Println(user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Code": 1,
			"Msg":  "用户名不存在",
		})
		return
	}
	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"Code": 1,
			"Msg":  "密码错误",
		})
		return
	}
	//设置cookie

	//
	c.JSON(http.StatusOK, gin.H{
		"Code": 0,
		"Msg":  "登录成功",
	})
}

//更新
func UpdateHandler(c *gin.Context) {
	//从请求中提取数据
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	fmt.Println(user)
	//检查用户id是否存在
	u, err := model.SelectUser(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
			"msg":   "未找到对应的id用户",
		})
		return
	} else {
		//检查名字和密码是否正确
		if u.Name != user.Name {
			c.JSON(http.StatusOK, gin.H{
				"error": "err",
				"msg":   "用户名错误",
			})
			return
		}
		if u.Password != user.Password {
			c.JSON(http.StatusOK, gin.H{
				"error": "err",
				"msg":   "密码错误",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "正确",
	})
	//对实例修改
	user.Password = c.Query("password")
	//将实例输入数据库修改
	err = model.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
			"msg":   "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

//注销
func DeleteHandler(c *gin.Context) {

}

//查询：通过名字
func FindName(c *gin.Context) {

}
