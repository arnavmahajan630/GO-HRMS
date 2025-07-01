package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mg MongoInstance

const (
	dbName = "hrms-db"
	dbUrl = "mongodb://localhost:27017"
)

func TimeOutCkeck() (context.Context, context.CancelFunc){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return ctx, cancel;
}

func Connect() error {
	ctx , cancel := TimeOutCkeck()
	defer cancel()
	client , err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if( err != nil) {return err}

	Mg = MongoInstance{
		Client: client,
		DB: client.Database(dbName),
	}

	return nil
}

