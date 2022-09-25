package api

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cache struct {
	db *mongo.Database
}

func MakeDBCache(db *mongo.Database) *Cache {
	return makeDBCache(db)
}

func makeDBCache(db *mongo.Database) *Cache {
	db.Collection("restaurants")
	db.Collection("restaurantsToSearch")
	return &Cache{db: db}
}

type storedRestaurant struct {
	ID                   int
	URI                  string
	RawRestaurantDetails RawRestaurantDetails
	RestaurantDetails    RestaurantDetails
	Empty                bool
}

//go:generate genopts --function SaveRestaurant verbose reallyVerbose
func (c *Cache) SaveRestaurant(ctx context.Context, uri string, raw RawRestaurantDetails, optss ...SaveRestaurantOption) error {
	opts := MakeSaveRestaurantOptions(optss...)
	log := makeLog("SaveRestaurant")

	id, name := raw.RestaurantProfile.Restaurant.RestaurantID, raw.RestaurantProfile.Restaurant.Name
	filter := bson.D{
		{Key: "id", Value: id},
	}
	if _, err := c.db.Collection("restaurants").DeleteMany(ctx, filter); err != nil {
		return err
	}
	rest, err := raw.Convert()
	if err != nil {
		return err
	}
	stored := storedRestaurant{
		URI:                  uri,
		RawRestaurantDetails: raw,
		RestaurantDetails:    rest.RestaurantDetails,
		Empty:                false,
	}
	if _, err := c.db.Collection("restaurants").InsertOne(ctx, stored); err != nil {
		return err
	}

	if opts.ReallyVerbose() {
		log.Printf("saved restaurant %q as %d", name, id)
	}
	return nil
}

func isEmptyResultError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "no documents in result")
}

type SaveRestaurantToSearchResult string

const SaveRestaurantToSearch_RestaurantAlreadySearched SaveRestaurantToSearchResult = "already searched"
const SaveRestaurantToSearch_RestaurantAlreadyWaiting SaveRestaurantToSearchResult = "already waiting"
const SaveRestaurantToSearch_Added SaveRestaurantToSearchResult = "added"
const SaveRestaurantToSearch_None SaveRestaurantToSearchResult = "none"

func (c *Cache) SaveRestaurantToSearch(ctx context.Context, restaurantURI string, optss ...SaveRestaurantOption) (SaveRestaurantToSearchResult, error) {
	opts := MakeSaveRestaurantOptions(optss...)
	log := makeLog("SaveRestaurantToSearch")

	filter := bson.D{
		{Key: "uri", Value: restaurantURI},
	}
	if res := c.db.Collection("restaurants").FindOne(ctx, filter); res.Err() == nil {
		if opts.ReallyVerbose() {
			log.Printf("restaurant already searched: %s", restaurantURI)
		}
		return SaveRestaurantToSearch_RestaurantAlreadySearched, nil
	} else if !isEmptyResultError(res.Err()) {
		return SaveRestaurantToSearch_None, res.Err()
	}
	if res := c.db.Collection("restaurantsToSearch").FindOne(ctx, filter); res.Err() == nil {
		if opts.ReallyVerbose() {
			log.Printf("restaurant waiting to be searched already: %s", restaurantURI)
		}
		return SaveRestaurantToSearch_RestaurantAlreadyWaiting, nil
	} else if !isEmptyResultError(res.Err()) {
		return SaveRestaurantToSearch_None, res.Err()
	}

	stored := storedRestaurant{
		URI:   restaurantURI,
		Empty: true,
	}
	if _, err := c.db.Collection("restaurantsToSearch").InsertOne(ctx, stored); err != nil {
		return SaveRestaurantToSearch_None, err
	}

	if opts.ReallyVerbose() {
		log.Printf("saved restaurant %s", restaurantURI)
	}
	return SaveRestaurantToSearch_Added, nil
}

func (c *Cache) DeleteRestaurantToSearch(ctx context.Context, restaurantURI string) error {
	filter := bson.D{
		{Key: "uri", Value: restaurantURI},
	}
	if _, err := c.db.Collection("restaurantsToSearch").DeleteMany(ctx, filter); err != nil {
		return err
	}
	return nil
}

type RestaurantToSearch struct {
	URI string
}

func (c *Cache) RestaurantsToSearch(ctx context.Context) ([]RestaurantToSearch, error) {
	cur, err := c.db.Collection("restaurantsToSearch").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var res []RestaurantToSearch
	for cur.Next(ctx) {
		var el storedRestaurant
		if err := cur.Decode(&el); err != nil {
			return nil, err
		}
		res = append(res, RestaurantToSearch{
			URI: el.URI,
		})
	}
	return res, nil
}
