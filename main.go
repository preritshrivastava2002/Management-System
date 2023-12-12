package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

type House struct {
	ID          string  `bson:"_id,omitempty"`
	Address     string  `bson:"address"`
	Description string  `bson:"description"`
	City        string  `bson:"city"`
	State       string  `bson:"state"`
	PostalCode  int     `bson:"postalCode"`
	Price       float64 `bson:"price"`
	ForSale     bool    `bson:"forSale"`
	Available   bool    `bson:"available"`
}

func main() {
	app := gofr.New()
	app.GET("/start", StartHandler)
	app.Start()
}

func connectToDb() *mongo.Collection {
	uri := os.Getenv("MONOGDB_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	database := client.Database("HomeManagmentSystem")
	collection := database.Collection("homes")
	return collection
}

func StartHandler(c *gofr.Context) (interface{}, error) {
	return "Welcome to the Project!", nil
}
