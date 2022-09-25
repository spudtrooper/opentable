package api

import (
	"context"
	"flag"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoDBName = flag.String("mongo_db_name", "opentable", "the name of the MongoDB database to use")
	mongoDBPort = flag.Int("mongo_db_port", 27017, "the port of the MongoDB database to use")
)

func ConnectToDB(ctx context.Context) (*mongo.Database, error) {
	port := *mongoDBPort
	dbName := *mongoDBName

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
