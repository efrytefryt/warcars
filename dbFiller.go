package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mongodb/mongo-go-driver/bson"
)

type dbFiller struct {
}

func (f dbFiller) addPlayers() {
	f.addPlayer()
	//	f.addPlayer()
	//	f.addPlayer()
}

func (dbFiller) addPlayer() {
	count, err := storedPlayers.EstimatedDocumentCount(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	p := &player{
		ID:   uuid.New().String(),
		Name: fmt.Sprintf("Player %v", count)}

	_, err = storedPlayers.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
}

func (dbFiller) readFromDB() {
	fmt.Println("Players in DB")

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

		fmt.Println("Player: " + jsonDoc)

		var p player
		err = json.Unmarshal([]byte(jsonDoc), &p)

		fmt.Println(p)
	}
	fmt.Println("Players in DB << END")
}
