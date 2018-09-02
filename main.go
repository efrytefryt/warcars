package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	client        *mongo.Client
	db            *mongo.Database
	storedPlayers *mongo.Collection
)

func init() {
	var err error
	client, err = mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	const dbName string = "mongoTest"
	db = client.Database(dbName)

	const playersCollName string = "players"
	storedPlayers = db.Collection(playersCollName)
}

func main() {

	//	dbFiller.addPlayers(dbFiller{})

	r := mux.NewRouter()
	r.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		playersHandler.post(playersHandler{}, w, r)
	}).Methods("POST")

	r.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		playersHandler.getAll(playersHandler{}, w, r)
	}).Methods("GET")

	r.HandleFunc("/players/{id}", func(w http.ResponseWriter, r *http.Request) {
		playersHandler.getOne(playersHandler{}, w, r)
	}).Methods("GET")

	r.HandleFunc("/auctionHouseItems", auctionHouseItemsListingHandler).Methods("GET")

	fmt.Println("working")
	dbFiller.readFromDB(dbFiller{})

	http.ListenAndServe(":80", r)

	fmt.Printf("Finished")
}
