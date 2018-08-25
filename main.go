package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type inventory struct {
	parts []string
}

type player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (p player) String() string {
	return fmt.Sprint("[Player| id:", p.ID, ", name:", p.Name, "]")
}

func playerAdditionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("playerAdditionHandler")
}

func playersListingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("playersListingHandler")
}

type auctionHouseItem struct {
	Name string `json:"myName"`
}

func auctionHouseItemsListingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auctionHouseItemsListingHandler")

	collection := db.Collection("qux")
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var items []auctionHouseItem

	for cur.Next(context.Background()) {
		doc := bson.NewDocument()
		if err := cur.Decode(doc); err != nil {
			log.Fatal(err)
		}

	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	client  *mongo.Client
	db      *mongo.Database
	players *mongo.Collection
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
	players = db.Collection(playersCollName)
}

func readFromDB() {

	cur, err := players.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		doc := bson.NewDocument()
		if err := cur.Decode(doc); err != nil {
			log.Fatal("Decode:", err)
		}

		jsonDoc := doc.ToExtJSON(true)

		var p player
		err = json.Unmarshal([]byte(jsonDoc), &p)

		fmt.Println(p)
	}
}

func addPlayers() {
	addPlayer()
	addPlayer()
	addPlayer()
}

func addPlayer() {
	count, err := players.EstimatedDocumentCount(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	p := &player{
		ID:   uuid.New().String(),
		Name: fmt.Sprintf("Player %v", count)}

	_, err = players.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	//	addPlayers()

	r := mux.NewRouter()
	r.HandleFunc("/players", playerAdditionHandler).Methods("POST")
	r.HandleFunc("/players", playersListingHandler).Methods("GET")
	r.HandleFunc("/auctionHouseItems", auctionHouseItemsListingHandler).Methods("GET")

	fmt.Println("working")
	readFromDB()

	http.ListenAndServe(":80", r)

	fmt.Printf("Finished")
}
