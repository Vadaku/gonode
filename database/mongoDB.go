package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Create client instance and check for any errors returned from "mongo.Connect()".
//If any errors then terminate the application using "panic()".
func ConnectToDB() {
	fmt.Println("Starting Mongodb on port 27017")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	//addToDB()

}

//Results struct is added to the database.
func AddToDB(r interface{}) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	usersCollection := client.Database("nevada").Collection("data")
	result, err := usersCollection.InsertOne(context.TODO(), r)
	fmt.Println("Mining result added to database", result.InsertedID)
	if err != nil {
		panic(err)
	}
}

//Lookup result according to rotation parameter.
func LookupInDB(rotation string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	usersCollection := client.Database("nevada").Collection("data")
	var result bson.M
	err = usersCollection.FindOne(context.TODO(), bson.M{"rotation": rotation}).Decode(&result)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
}
