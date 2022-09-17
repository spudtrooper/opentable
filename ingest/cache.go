package ingest

import (
	"context"
	"log"

	"github.com/spudtrooper/opentable/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cache interface {
	SaveRestaurant(ctx context.Context, rest api.RawRestaurantDetails, optss ...SaveRestaurantOption) error
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
	RestaurantDetails    api.RestaurantDetails
}

//go:generate genopts --function SaveRestaurant verbose
func (c *dbCache) SaveRestaurant(ctx context.Context, raw api.RawRestaurantDetails, optss ...SaveRestaurantOption) error {
	opts := MakeSaveRestaurantOptions(optss...)

	id, name := raw.RestaurantProfile.Restaurant.RestaurantID, raw.RestaurantProfile.Restaurant.Name
	filter := bson.D{
		{Key: "id", Value: id},
		{Key: "name", Value: name},
	}
	if _, err := c.db.Collection("restaurants").DeleteMany(ctx, filter); err != nil {
		return err
	}
	rest, err := raw.Convert()
	if err != nil {
		return err
	}
	stored := storedRestaurant{
		RawRestaurantDetails: raw,
		RestaurantDetails:    rest.RestaurantDetails,
	}
	if _, err := c.db.Collection("restaurants").InsertOne(ctx, stored); err != nil {
		return err
	}

	if opts.Verbose() {
		log.Printf("saved restaurant %q as %d", name, id)
	}
	return nil
}
