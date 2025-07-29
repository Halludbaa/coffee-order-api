package v1

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"coffee/internal/model/apperrors"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SessionRepo struct {
	conn *sqlx.DB
	log 	*logrus.Logger
}

func NewSessionRepo(conn *sqlx.DB, log *logrus.Logger) model.SessionRepo {
	return &SessionRepo{
		conn, log,
	}
}

func (sRepo *SessionRepo) Store(ctx context.Context, request *entity.Session) (error) {
	query := `INSERT INTO sessions_manager (user_id, user_name, user_agent, token) VALUES (:user_id, :user_name, :user_agent, :token)`

	_, err := sRepo.conn.NamedQueryContext(ctx, query, request)
	if err != nil {
		return apperrors.NewInternal()
	}

	return nil
}

func (sRepo *SessionRepo) Remove(ctx context.Context, request *entity.Session) (error) {
	query := `DELETE FROM sessions_manager WHERE token = :token`
	
	_, err := sRepo.conn.NamedQueryContext(ctx, query, request)
	if err != nil {
		return apperrors.NewInternal()
	}

	return nil
}

func (sRepo *SessionRepo) FindByUserId(ctx context.Context,  record *entity.Session) (error) {
	query := `SELECT user_name, token FROM sessions_manager WHERE user_id := $1`

	if err := sRepo.conn.GetContext(ctx, record, query, record.UserID); err != nil {
		return apperrors.NewInternal()
	}

	return nil	
}

func (sRepo *SessionRepo) FindByToken(ctx context.Context, record *entity.Session) (error) {
	query := `SELECT user_name, token FROM sessions_manager WHERE token b= $1`

	if err := sRepo.conn.GetContext(ctx, record, query, record.Token); err != nil {
		sRepo.log.Warn(err)
		return apperrors.NewAuthorization("you already logout!")
	}

	return nil	
}
