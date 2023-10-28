package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var databaseName = "githubIssuesNotificator"

var client *mongo.Client
var ctx = context.Background()

func Database() *mongo.Database {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017/githubIssuesNotificator?authSource=admin")
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalln(err)
	}

	return client.Database(databaseName)
}

func Close() error {
	return client.Disconnect(ctx)
}
