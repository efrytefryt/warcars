package main

import (
	"fmt"
	"net/http"
)

type auctionHouseItem struct {
	Name string `json:"myName"`
}

func auctionHouseItemsListingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auctionHouseItemsListingHandler")
}
