package controller

//定义一个数据类型，名字叫ResCode，特征和int64相同
type ResCode int64

const ( //摄制常量
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeMysql
)

//定义一个哈希表
var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",

	CodeMysql: "mysql数据库操作有误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]

	if !ok { //找不到对应c值时
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
