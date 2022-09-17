package ingest

import (
	"context"

	"github.com/spudtrooper/opentable/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cache interface {
	SaveRestaurant(ctx context.Context, rest api.RawRestaurantDetails) error
	// GetRestaurant(ctx context.Context, query string) (*api.RawRestaurantDetails, error)
}

type dbCache struct {
	db *mongo.Database
}

func MakeDBCache(db *mongo.Database) Cache {
	return makeDBCache(db)
}

func makeDBCache(db *mongo.Database) *dbCache {
	db.Collection("restaurants")
	return &dbCache{db: db}
}

type storedRestaurant struct {
	RawRestaurantDetails api.RawRestaurantDetails
}

func (c *dbCache) SaveRestaurant(ctx context.Context, rest api.RawRestaurantDetails) error {
	filter := bson.D{
		{Key: "id", Value: rest.RestaurantProfile.Restaurant.RestaurantID},
		{Key: "name", Value: rest.RestaurantProfile.Restaurant.Name},
	}
	if _, err := c.db.Collection("restaurants").DeleteMany(ctx, filter); err != nil {
		return err
	}
	stored := storedRestaurant{
		RawRestaurantDetails: rest,
	}
	if _, err := c.db.Collection("restaurants").InsertOne(ctx, stored); err != nil {
		return err
	}
	return nil
}
