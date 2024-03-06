package repositories

import (
	"context"
	"math/rand"
	"test-go/core/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repoPower struct {
	db         *mongo.Database
	collection string
}

func NewPowerRepository(db *mongo.Database, collection string) *repoPower {
	return &repoPower{
		db:         db,
		collection: collection,
	}
}

func (r repoPower) CreatePower() error {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	for i := 0; i < 1000; i++ {
		randomInput := rand.Intn(1000) + 1
		randomActive := rand.Intn(1000) + 1

		payload := models.PowerRepository{
			ActivePower: randomActive,
			PowerInput:  randomInput,
		}

		_, err := r.db.Collection(r.collection).InsertOne(ctx, payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r repoPower) GetPower() (result []models.PowerRepository, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := bson.D{}
	res, err := r.db.Collection(r.collection).Find(ctx, query)
	if err != nil {
		return result, err
	}
	err = res.All(ctx, &result)
	return result, err
}
