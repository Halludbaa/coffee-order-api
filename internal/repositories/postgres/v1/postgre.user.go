package v1

import (
	"coffee/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	conn *sqlx.DB
	log 	*logrus.Logger
}

func NewUserRepo(conn *sqlx.DB, log *logrus.Logger) model.UserRepository {
	return &UserRepo{
		conn: conn,
		log: log,
	}
}
