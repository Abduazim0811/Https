package storage

import (
	"Items/internal/https/api/handler"
	"Items/internal/infrastructura/repository/mongodb"
	"Items/internal/service"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Storage() (*mongo.Client, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("Items").Collection("items")
	return client, collection, nil
}

func Handler() *handler.ItemsHandler{
	client, collection, err := Storage()
	if err != nil {
		log.Println("connection mongodb error")
		return nil
	}

	repo := mongodb.NewItemsMongodb(client, collection)

	service := service.NewItems(repo)

	handler := handler.NewService(service)

	return handler
}
