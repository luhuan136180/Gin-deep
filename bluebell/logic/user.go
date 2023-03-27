package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"fmt"
)

//存放业务逻辑的处理
//各种操作的调用（大杂烩）
func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		fmt.Println(err)
		return err
	}
	//2.生成uid
	userID := snowflake.GenID()
	//构造user实例
	user := &models.User{
		UserID:   userID,
		Password: p.Password,
		Name:     p.Username,
	}

	//4.保存进数据库
	return mysql.InsertUser(user)

}

func Login(p *models.ParamLogin) (token string, err error) {
	//(检查用户存不存在)
	user := &models.User{
		Name:     p.Username,
		Password: p.Password,
	}

	//传递的是指针，
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	//生成jwt
	return jwt.GenToken(user.UserID, user.Name)

}
