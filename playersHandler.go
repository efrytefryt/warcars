package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson"
)

type playersHandler struct {
}

func (playersHandler) post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Body:", body)

	var p player
	json.Unmarshal(body, &p)
	fmt.Println(p)

	res, err := storedPlayers.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted:", res.InsertedID)
}

func (playersHandler) get(w http.ResponseWriter, r *http.Request) {
	var players []player

	cur, err := storedPlayers.Find(context.Background(), nil)
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
		if err = json.Unmarshal([]byte(jsonDoc), &p); err != nil {
			log.Fatal("Unmarshal:", err)
		}

		players = append(players, p)
	}

	if err = json.NewEncoder(w).Encode(players); err != nil {
		log.Fatal("Encode:", err)
	}
}
