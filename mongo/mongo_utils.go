package mongo

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"os"
)

func loadConnection() *mongo.Client {
	client, err := mongo.NewClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetCollection(dbname, cname string) *mongo.Collection {
	client := loadConnection()

	return client.Database(dbname).Collection(cname)
}
