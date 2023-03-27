package models

type User struct {
	UserID   int64  `db:"user_id,string"`
	Name     string `db:"username"`
	Password string `db:"password"`
}
