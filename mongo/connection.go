package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToMongo() {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	_, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to mongod")
}
