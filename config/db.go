package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnect() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB server")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("tumorBoard").Collection(collectionName)
}

func AddIndex(client *mongo.Client, collection string, indexKeys interface{}) (*mongo.Collection, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Minute)
	serviceCollection := client.Database("tumorBoard").Collection(collection)
	_, err := serviceCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: indexKeys,
	})
	if err != nil {
		return nil, err
	}
	serviceCollectionWithIndex := client.Database("tumorBoard").Collection(collection)
	return serviceCollectionWithIndex, nil
}
