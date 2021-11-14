package DB

import (
	"admin/Util"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	url        string
	name       string
	collection string
	user       string
	pwd        string
}

type MongoInstance struct {
	RpyClient *mongo.Client
	RpyDBCol  *mongo.Collection
}

var MI MongoInstance


func MongoConnect() {
	url, name, c := Util.GoDotEnvVariable("DB_URL"), Util.GoDotEnvVariable("DB_NAME"), Util.GoDotEnvVariable("DB_COLLECTION")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("Cannot connect database")
	}
	logs := client.Database(name).Collection(c)
	count, err := logs.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)


	fmt.Println("DB connected!")
	MI = MongoInstance{
		RpyClient: client,
		RpyDBCol:  logs,
	}
}
