package resources

import (
	"context"
	"fmt"
	"github.com/narozhniyy/test/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
	"time"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through GetMongoClient().*/
var client *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

// Get mongo db client
func getMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_HOST"))
		c, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientError = err
		}

		err = c.Ping(context.TODO(), nil)
		if err != nil {
			clientError = err
		}

		client = c
	})

	return client, clientError
}

// Insert document to mongo db
func InsertDocument(t *models.Table) (*mongo.InsertOneResult, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(os.Getenv("MONGODB_DB")).Collection(os.Getenv("MONGODB_COLLECTION"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, t)
	if err != nil {
		return nil, fmt.Errorf("execute query error: %v", err)
	}

	return result, nil
}

// Get document from dynamo db
func GetDocument(table int64) (*models.Table, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(os.Getenv("MONGODB_DB")).Collection(os.Getenv("MONGODB_COLLECTION"))

	var ep *models.Table
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"table": table, "active": true}, options.Find().SetLimit(1))
	if err != nil {
		return nil, fmt.Errorf("execute query error: %v", err)
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&ep)
		if err != nil {
			return nil, fmt.Errorf("decode error: %v", err)
		}
	}

	return ep, nil
}

// Update document in mongo db
func UpdateDocument(table *models.Table, tn int64) (*models.Table, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(os.Getenv("MONGODB_DB")).Collection(os.Getenv("MONGODB_COLLECTION"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"$and": []interface{}{bson.M{"table": tn}, bson.M{"active": true}}}
	update := bson.M{"$set": bson.M{"table": table.Table, "active": table.Active, "guests": table.Guests}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("execute query error: %v", err)
	}

	return table, nil
}
