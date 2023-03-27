package logic

import (
	"bluebell0002/dao/mysql"
	"bluebell0002/models"
	"bluebell0002/pkg/jwt"
	"bluebell0002/pkg/snowFlake"
	"fmt"
)

//存放业务逻辑的处理
//各种操作的调用（大杂烩）
func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		fmt.Println(err)
		return
	}
	//2.生成uid
	useID := snowFlake.GenID()
	//构造user实例
	user := &models.User{
		UserID:   useID,
		Password: p.Password,
		Name:     p.Username,
	}
	//4.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	//
	user := &models.User{
		Name:     p.Username,
		Password: p.Password,
	}

	//传递指针
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Name)
}
