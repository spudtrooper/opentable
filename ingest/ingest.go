package ingest

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(ctx context.Context) (*mongo.Database, error) {
	const port = 27017
	const dbName = "opentable"

	uri := fmt.Sprintf("mongodb://localhost:%d", port)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Printf("connected to %q @ %s", dbName, uri)

	return client.Database(dbName), nil
}
