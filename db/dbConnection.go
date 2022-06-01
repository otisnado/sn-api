package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUser string = os.Getenv("MONGO_USER")
var mongoPassword string = os.Getenv("MONGO_PASSWORD")
var mongoHost string = os.Getenv("MONGO_HOST")

/*MongoCN objeto de conexion a base de datos*/
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://" + mongoUser + ":" + mongoPassword + "@" + mongoHost + "/?retryWrites=true&w=majority")

/*ConnectDB conexion a base de datos */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la DB")
	return client
}

/*CheckConnection ping a base de datos */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
