package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	KodeProduk string `bson:"kodeProduk,omitempty" json:"kodeProduk,omitempty"`
	Kuantitas  int    `bson:"kuantitas,omitempty" json:"kuantitas,omitempty"`
}

// Request Model
type ModelRequest struct {
	KodeProduk string `json:"kodeProduk,omitempty"`
	Kuantitas  int    `json:"kuantitas,omitempty"`
}

// Response Model
type ModelResponse struct {
	ResponseMessage string `json:"responseMessage"`
}

type ModelResponse2 struct {
	ResponseMessage string `json:"responseMessage"`

	KodeProduk string `json:"kodeProduk,omitempty"`
	Kuantitas  int    `json:"kuantitas,omitempty"`
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

	fmt.Printf("Successfully connected and pinged.\n")

}

func initServer() {
	log.Printf("Server listener started.\n\n")

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
	var res *mongo.UpdateResult
	var update primitive.D
	var upsertedId interface{}
	var responseJson []byte
	var modelResponse ModelResponse
	var modelProduct ModelProduct

	// Parse body
	var modelRequest ModelRequest
	err = json.NewDecoder(r.Body).Decode(&modelRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Kode Product : %v\n", modelRequest.KodeProduk)
	fmt.Printf("Kuantitas : %v\n", modelRequest.Kuantitas)

	fmt.Printf("Request insert document\n")

	// Create structure
	modelProduct = ModelProduct{
		KodeProduk: modelRequest.KodeProduk,
		Kuantitas:  modelRequest.Kuantitas,
	}

	// Insert to database
	coll = client.Database(databaseName).Collection(collectionName)
	filter = bson.D{primitive.E{Key: "kodeProduk", Value: modelProduct.KodeProduk}}
	opts = options.Update().SetUpsert(true)
	update = bson.D{{Key: "$set", Value: modelProduct}}
	res, err = coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Println(err)
		return
	}
	upsertedId = res.UpsertedID

	fmt.Printf("Document UpsertedID with id : %v\n\n", upsertedId)

	if upsertedId != nil {
		// Create response
		modelResponse = ModelResponse{
			ResponseMessage: "Product data inserted to database",
		}
	} else {
		// Create response
		modelResponse = ModelResponse{
			ResponseMessage: "Product data updated to database",
		}

	}

	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		log.Println(err)
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
	var modelResponse ModelResponse
	var modelResponse2 ModelResponse2
	var filter primitive.D

	// Parse body
	queryParameter = r.URL.Query()
	productId = queryParameter["kodeProduk"][0]

	fmt.Printf("Request read document by productId : %v\n", productId)

	// begin find
	myCollection = client.Database(databaseName).Collection(collectionName)

	filter = bson.D{{Key: "kodeProduk", Value: productId}}

	var modelProduct ModelProduct
	err = myCollection.FindOne(
		context.TODO(),
		filter,
	).Decode(&modelProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			modelResponse = ModelResponse{
				ResponseMessage: "Data not available",
			}

			responseJson, err = json.Marshal(modelResponse)
			if err != nil {
				log.Println(err)
				return
			}

			w.Header().Set("content-type", "application/json")
			w.Write(responseJson)

			log.Println(err)

			return
		}
		log.Println(err)
		return
	}
	// end find

	fmt.Printf("Document finded\n\n")

	// Create response
	modelResponse2 = ModelResponse2{
		ResponseMessage: "Product data available in database",
		KodeProduk:      modelProduct.KodeProduk,
		Kuantitas:       modelProduct.Kuantitas,
	}

	responseJson, err = json.Marshal(modelResponse2)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	var queryParameter url.Values
	var productId string
	var coll *mongo.Collection
	var err error
	var filter primitive.D
	var result *mongo.DeleteResult
	var modelResponse ModelResponse
	var responseJson []byte

	// Parse body
	queryParameter = r.URL.Query()
	productId = queryParameter["kodeProduk"][0]

	fmt.Printf("Request delete document by productId : %v\n", productId)

	coll = client.Database(databaseName).Collection(collectionName)
	filter = bson.D{{Key: "kodeProduk", Value: productId}}
	result, err = coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return
	}

	if result.DeletedCount == 0 {
		modelResponse = ModelResponse{
			ResponseMessage: "Data not available",
		}

		responseJson, err = json.Marshal(modelResponse)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(responseJson)
		log.Println("Can't find your data")
		return
	}

	fmt.Printf("Document deleted\n\n")

	// Create response
	modelResponse = ModelResponse{
		ResponseMessage: "Product data deleted from database",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}
