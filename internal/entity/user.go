package entity

type User struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Fullname string `db:"fullname"`
	Password string `db:"password"`
	Email    string `db:"email"`
}