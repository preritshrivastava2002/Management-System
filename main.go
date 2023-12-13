package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
	app.GET("/homes", HomesHandler)
	app.POST("/homes", CreateHomeHandler)
	app.GET("/homes/:id", GetHomeHandler)
	app.Start()
}

func connectToDb() *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error Loading the file ", err)
	}
	uri := os.Getenv("MONOGDB_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("HomeManagmentSystem")
	collection := database.Collection("homes")
	return collection
}

func StartHandler(c *gofr.Context) (interface{}, error) {
	return "Welcome to the Project!", nil
}

func HomesHandler(c *gofr.Context) (interface{}, error) {
	collection := connectToDb()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var homes []House
	if err := cursor.All(ctx, &homes); err != nil {
		return nil, err
	}

	return homes, nil
}

func CreateHomeHandler(c *gofr.Context) (interface{}, error) {
	collection := connectToDb()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newHome House

	res, err := collection.InsertOne(ctx, newHome)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func GetHomeHandler(c *gofr.Context) (interface{}, error) {
	collection := connectToDb()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Param("id")
	fmt.Println(id)
	var home House
	err := collection.FindOne(ctx, bson.M{"_id": (id)}).Decode(&home)
	if err != nil {
		return nil, err
	}

	return home, nil
}
