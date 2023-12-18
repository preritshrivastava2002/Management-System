package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/preritshrivastava2002/Management-System/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

func connectToDb() *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("YOUR MONGODB_URL"))
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

func AllHomeHandler(c *gofr.Context) (interface{}, error) {
	collection := connectToDb()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var homes []model.House
	er := cursor.All(ctx, &homes)
	if er != nil {
		return nil, er
	}

	return homes, nil
}

func CreateHomeHandler(c *gofr.Context) (interface{}, error) {
	collection := connectToDb()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newHome model.House
	err := json.NewDecoder(c.Request().Body).Decode(&newHome)
	if err != nil {
		return nil, err
	}
	fmt.Println(newHome)
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

	id := c.PathParam("id")

	var home model.House
	parsedId, _ := primitive.ObjectIDFromHex(id)
	fmt.Println("parsed: ", parsedId)

	err := collection.FindOne(ctx, bson.M{"_id": (parsedId)}).Decode(&home)
	if err != nil {
		return nil, err
	}

	return home, nil
}

func UpdateHomeHandler(c *gofr.Context) (interface{}, error) {
	collection := connectToDb()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var updatedHome model.House
	err := json.NewDecoder(c.Request().Body).Decode(&updatedHome)
	if err != nil {
		return nil, err
	}

	id := c.PathParam("id")
	parseId, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{"$set": updatedHome}
	_, er := collection.UpdateOne(ctx, bson.M{"_id": parseId}, update)
	if er != nil {
		return nil, er
	}

	return "House Updated Successfully", nil
}

func DeleteHomeHandler(c *gofr.Context) (interface{}, error) {
	collection := connectToDb()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.PathParam("id")
	parseID, _ := primitive.ObjectIDFromHex(id)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": parseID})
	if err != nil {
		return nil, err
	}

	return "Deleted Successfully House", nil
}
