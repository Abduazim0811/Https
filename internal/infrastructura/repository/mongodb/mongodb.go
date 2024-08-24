package mongodb

import (
	"Items/internal/entity/items"
	"Items/internal/infrastructura/repository"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemsMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewItemsMongodb(client *mongo.Client, collection *mongo.Collection) repository.ItemsRepository {
	return &ItemsMongodb{client: client, collection: collection}
}

func (i *ItemsMongodb) AddItems(req items.CreateItems) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := i.collection.InsertOne(ctx, req)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (i *ItemsMongodb) GetbyidItems(id string) (*items.Items, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID format:", err)
		return nil, err
	}

	var item items.Items
	err = i.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("item not found")
		}
		log.Println("Failed to get item:", err)
		return nil, err
	}
	return &item, nil
}


func (i *ItemsMongodb) GetAll() ([]*items.Items, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := i.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Failed to get all items:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var itemsList []*items.Items
	for cursor.Next(ctx) {
		var item items.Items
		if err := cursor.Decode(&item); err != nil {
			log.Println("Failed to decode item:", err)
			return nil, err
		}
		itemsList = append(itemsList, &item)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return itemsList, nil
}

func (i *ItemsMongodb) Update(id string, updateData items.Items) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID format:", err)
		return err
	}

	update := bson.M{
		"$set": updateData,
	}

	_, err = i.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Println("Failed to update item:", err)
		return err
	}

	return nil
}

func (i *ItemsMongodb) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID format:", err)
		return err
	}

	_, err = i.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Println("Failed to delete item:", err)
		return err
	}

	return nil
}
