package mysql

import (
	"bluebell/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var Db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {

	//Sprintf()根据格式说明符格式化并返回结果字符串。
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	Db, err = sqlx.Connect("mysql", dns)

	if err != nil {
		zap.L().Error("connect DB failed,err:", zap.Error(err))
		return
	}

	Db.SetMaxOpenConns(cfg.MaxOpenConns)
	Db.SetMaxIdleConns(cfg.MaxIdleConns)

	return
}
