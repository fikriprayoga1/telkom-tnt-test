package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Variable Block
var client *mongo.Client

const databaseName string = "telkom"

const collectionName string = "product"

const port string = "8080"

// Object Block
// Database Model
type ModelProduct struct {
	Message    string `bson:"message,omitempty" json:"message,omitempty"`
	KodeProduk string `bson:"kodeProduk,omitempty" json:"kodeProduk,omitempty"`
	Kuantitas  int    `bson:"kuantitas,omitempty" json:"kuantitas,omitempty"`
}

func main() {
	initDatabase()
	initServer()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func initDatabase() {
	// Replace the uri string with your MongoDB deployment's connection string.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	databaseUrl := "mongodb://" + os.Getenv("MONGO_HOST") + ":27017"
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	err2 := client.Ping(ctx, readpref.Primary())
	if err2 != nil {
		panic(err2)
	}

	log.Println("Successfully connected and pinged")

}

func initServer() {
	log.Println("Server listener started")

	http.HandleFunc("/product/set", setProduct)
	http.HandleFunc("/product/read", readProduct)
	http.HandleFunc("/product/delete", deleteProduct)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setProduct(w http.ResponseWriter, r *http.Request) {
	var err error
	var coll *mongo.Collection
	var filter primitive.D
	var opts *options.UpdateOptions
	var update primitive.D
	var responseJson []byte
	var modelProduct ModelProduct

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelProduct)
	if err != nil {
		errorHandler(w, "Body request is wrong", err)
		return
	}

	// Insert to database
	coll = client.Database(databaseName).Collection(collectionName)
	filter = bson.D{primitive.E{Key: "kodeProduk", Value: modelProduct.KodeProduk}}
	opts = options.Update().SetUpsert(true)
	update = bson.D{{Key: "$set", Value: modelProduct}}
	_, err = coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		errorHandler(w, "coll.UpdateOne error", err)
		return
	}

	modelProduct = ModelProduct{
		Message: "Product data has set to database",
	}

	responseJson, err = json.Marshal(modelProduct)
	if err != nil {
		errorHandler(w, "Response parse to json failed", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func readProduct(w http.ResponseWriter, r *http.Request) {
	var queryParameter url.Values
	var productId string
	var err error
	var myCollection *mongo.Collection
	var responseJson []byte
	var modelResponse ModelProduct
	var filter primitive.D

	// Parse body
	queryParameter = r.URL.Query()
	productId = queryParameter["kodeProduk"][0]

	// begin find
	myCollection = client.Database(databaseName).Collection(collectionName)

	filter = bson.D{{Key: "kodeProduk", Value: productId}}

	var modelDatabaseProduct ModelProduct
	err = myCollection.FindOne(
		context.TODO(),
		filter,
	).Decode(&modelDatabaseProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorHandler(w, "Data not available", err)
			return
		}
		errorHandler(w, "Decode error", err)
		return
	}
	// end find

	// Create response
	modelResponse = ModelProduct{
		Message:    "Product data available in database",
		KodeProduk: modelDatabaseProduct.KodeProduk,
		Kuantitas:  modelDatabaseProduct.Kuantitas,
	}

	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		errorHandler(w, "Parsing JSON error", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	var coll *mongo.Collection
	var err error
	var filter primitive.D
	var result *mongo.DeleteResult
	var modelProduct ModelProduct
	var responseJson []byte

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelProduct)
	if err != nil {
		errorHandler(w, "Body request is wrong", err)
		return
	}

	coll = client.Database(databaseName).Collection(collectionName)
	filter = bson.D{{Key: "kodeProduk", Value: modelProduct.KodeProduk}}
	result, err = coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		errorHandler(w, "Database error", err)
		return
	}

	if result.DeletedCount == 0 {

		errorHandler(w, "Data not available", err)
		return
	}

	// Create response
	modelProduct = ModelProduct{
		Message: "Product data deleted from database",
	}
	responseJson, err = json.Marshal(modelProduct)
	if err != nil {
		errorHandler(w, "Parsing JSON error", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func errorHandler(w http.ResponseWriter, message string, err error) {
	var responseJson []byte

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	modelProduct := ModelProduct{
		Message: message,
	}

	responseJson, err = json.Marshal(modelProduct)
	if err != nil {
		w.Write([]byte("Parsing error"))
		return
	}

	w.Write(responseJson)

}
