package database

import (
	"context"
	"github.com/K-shir0/ajisai-api-server/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type WeathersRepository struct {
	*mongo.Collection
}

func (repo *WeathersRepository) FindAll() ([]bson.M, error) {

	o := options.Find()

	o.SetSort(bson.D{{"_id", -1}})
	o.SetLimit(100)

	cur, err := repo.Find(context.TODO(), bson.M{}, o)
	if err != nil {
		return nil, err
	}

	var weathers []bson.M
	if err = cur.All(context.TODO(), &weathers); err != nil {
		return nil, err
	}

	return weathers, nil
}

func (repo *WeathersRepository) Store(w *domain.Weather) (*mongo.InsertOneResult, error) {
	// set date
	now := time.Now()
	w.UpdatedAt, w.CreatedAt = now, now

	return repo.InsertOne(context.TODO(), w)
}
