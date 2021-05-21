package infrastructure

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type SqlHandler struct {
	Conn *mongo.Database
}

func (r *Router) NewSqlHandler() *mongo.Database {
	uri := "mongodb://" + r.config.MongoDB.User + ":" + r.config.MongoDB.Password + "@" + r.config.MongoDB.Host + ":" + r.config.MongoDB.Port
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	return client.Database(r.config.MongoDB.Database)
}

func Execute() {

}
