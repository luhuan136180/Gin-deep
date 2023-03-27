package controller

import (
	"bluebell0002/dao/mysql"
	"bluebell0002/logic"
	"bluebell0002/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//登录的业务的入口
func SignUpHandler(c *gin.Context) {
	//1.参数校验(controller层)
	p := new(models.ParamSignUp)
	if err := c.ShouldBind(p); err != nil {
		//请求参数有误，直接返回请求
		//记录日志
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	//2.业务处理(server/logic 层处理)
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseErrorWithMsg(c, CodeMysql, err.Error())
		return
	}
	//3.返回相应_成功
	ResponseSuccess(c, "注册成功")
}

func LoginHandler(c *gin.Context) {
	//从前端获取登录需要的信息（用户名，用户密码）,对参数进行校验
	//创建实例接收数据
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return

		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
	}

	//业务处理
	token, err := logic.Login(p)

	if err != nil {
		zap.L().Error("Logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeMysql, err.Error())
		return
	}
	ResponseSuccess(c, token)
}
