package v1

import (
	"context"
	"todo_api/internal/entity"
	"todo_api/internal/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VtuberRepo struct {
	// mconn *mongo.Client
	log *logrus.Logger
	collection *mongo.Collection
}

func NewVtuberRepo(mconn *mongo.Client, log *logrus.Logger, db string, collectName string) model.VtuberRepo {
	collection := mconn.Database(db).Collection(collectName)
	return &VtuberRepo{
		log, collection,
	}
}

func (repo *VtuberRepo) FindAll(ctx context.Context) (*[]entity.Vtuber, error) {
	var result []entity.Vtuber
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		repo.log.Warn(err)
		return nil, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		repo.log.Warn(err)
		return nil, err
	}

	return &result, nil
}