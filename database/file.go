package database

import (
	"fmt"
	"log"
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ashurai/fileUploader/model"
	"go.mongodb.org/mongo-driver/bson"
)

func connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	//fmt.Println("%T", client)
	return client
}

// SaveFile is a function to save files in DB
func SaveFile(file *model.File) (interface{}, error) {
	db := connect()
	coll := db.Database("fileManager").Collection("files")
	result, err := coll.InsertOne(context.TODO(), file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err
}

// GetFiles list out all the records from collection
func GetFiles(offset, limit int) []*model.Files {
	db := connect()
	collection := db.Database("fileManager").Collection("files")
	var results []*model.Files


	findOptions := options.Find()
	findOptions.SetLimit(int64(limit)).SetSkip(int64(offset))
	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem model.Files
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}

