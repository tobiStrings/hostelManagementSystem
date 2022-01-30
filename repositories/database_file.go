package repositories

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func EnvMongoURI() string {

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	err:=os.Setenv("MONGODB_URL", "mongodb://localhost:27017/hostel_management")
	if err!=nil {
		log.Fatal("error getting mongo instance")
	}
	err=os.Setenv("PORT", "8080")
	if err!=nil {
		log.Fatal("error getting port")
	}

	return os.Getenv("MONGODB_URL")
}

func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection {

	collection := client.Database("hostel_management").Collection(collectionName)

	return collection
}

func DbInstance() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {

		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DbInstance()