package mysql

import (
	"bluebell0002/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

//把每一步数据库操作都封装进dao层
//待logic层根据业务需求调用

const secret = "mhx.com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err := DB.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密

	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = DB.Exec(sqlStr, user.UserID, user.Name, user.Password)
	return

}

//encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select user_id,username,password from user where username=?"
	err = DB.Get(user, sqlStr, user.Name)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	err = DB.Get(user, sqlStr, uid)
	return
}
