package config

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewAppDatabase() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("app.database.uri")))
	if err != nil {
		panic(err)
	}
	// # Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	return client.Database(viper.GetString("app.database.dbname"))
}
