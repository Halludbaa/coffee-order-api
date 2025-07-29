package entity

type Session struct {
	UserID    string `db:"user_id"`
	Username  string `db:"user_name"`
	UserAgent string `db:"user_agent"`
	Token     string `db:"token"`
}