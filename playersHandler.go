package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
)

type playersHandler struct {
}

func (playersHandler) post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var p player
	json.Unmarshal(body, &p)

	res, err := storedPlayers.InsertOne(context.Background(), p)
	if err != nil {
		w.WriteHeader(400)
		log.Println("Player already exists:\n", res.InsertedID)
	} else {
		log.Println("Player inserted:", res.InsertedID)
	}
}

func (playersHandler) getAll(w http.ResponseWriter, r *http.Request) {
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

func (playersHandler) getOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

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

		if *p.ID == id {
			if err = json.NewEncoder(w).Encode(p); err != nil {
				log.Fatal("Encode:", err)
			}
			return
		}
	}

	log.Println("Player not found:", id)
	w.WriteHeader(400)
}
