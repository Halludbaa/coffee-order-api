package v1

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"coffee/internal/model/converter"
	"context"

	"github.com/gofiber/fiber/v2"
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

func (r *UserRepo) Store(ctx context.Context, request *model.SignUpRequest) (*entity.User, error) {
	userRecord := converter.SignUpToUser(request)
	r.log.Debug(userRecord)
	query := `INSERT INTO users (username, fullname, email, password) VALUES (:username, :fullname, :email, :password) RETURNING *`

	rows, err := r.conn.NamedQueryContext(ctx, query, userRecord)
	if err != nil {
		return nil, err
	}
	inserted := new(entity.User)

	defer rows.Close()

	if rows.Next(){
		if err := rows.StructScan(inserted); err != nil {
			r.log.WithError(err).Warn("can't scan rows")
			return nil, err
		}
		r.log.Debug(*inserted)
		return inserted, nil
	}
	return nil, fiber.ErrInternalServerError


}

func (r *UserRepo) Remove(ctx context.Context, username string) (error) {
	return nil
}

func (r *UserRepo) Update(ctx context.Context, userUpdateReq *model.UserUpdateRequest) (error) {
	return nil
}

func (r *UserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	userRecord := new(entity.User)

	query := `SELECT id, username, fullname, email, password FROM users WHERE username = $1`
	if err := r.conn.GetContext(ctx, userRecord, query, username); err != nil {
		return &entity.User{}, err
	}

	return userRecord, nil
}

func (r *UserRepo) FindById(ctx context.Context, Id string) (*entity.User, error) {
	return &entity.User{}, nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return &entity.User{}, nil
}
