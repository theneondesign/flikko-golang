package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func flikko() {
	allItems()
	handleRoutes()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homePage")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	// params := mux.Vars(r)
	// fmt.Println(params["ID"])
	itemID := mux.Vars(r)
	flag := false
	for i := 0; i < len(items); i++ {
		if itemID["ID"] == items[i].ID {
			json.NewEncoder(w).Encode(items[i])
			flag = true
			break
		}
	}

	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Items
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = strconv.Itoa(rand.Intn(1000))
	currentTime := time.Now().Format("02-01-2006 15:04")
	item.ProductPostedTime = currentTime
	item.ProductPostedFriendlyTime = rand.Intn(59)
	item.ProductLikeCount = rand.Intn(100)
	item.ProductViewCount = rand.Intn(100)
	// item.ProductQuantity = rand.Intn(10)
	item.ProductQuantityUnits = "Units"
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, product := range items {
		if product.ID == params["ID"] {
			items = append(items[:index], items[index+1:]...)
			var item Items
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.ID = params["ID"]
			items = append(items, item)
			flag = true
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, product := range items {
		if product.ID == params["ID"] {
			items = append(items[:index], items[index+1:]...)
			flag = true
			json.NewEncoder(w).Encode(map[string]string{"status": "Success"})
			return
		}
	}

	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/getitems", getItems).Methods("GET")
	router.HandleFunc("/getitem/{ID}", getItem).Methods("GET")
	router.HandleFunc("/create", createItem).Methods("POST")
	router.HandleFunc("/delete/", deleteItem).Queries("id", "{id}").Methods("DELETE")
	router.HandleFunc("/update/", updateItem).Queries("id", "{id}").Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082", router))
}

type Items struct {
	ID                        string `json:"ID"`
	ProductName               string `json:"product_name"`
	ProductDescription        string `json:"product_description"`
	ProductType               string `json:"product_type"`
	ProductCategory           string `json:"product_category"`
	ProductLikeCount          int    `json:"product_like_count"`
	ProductViewCount          int    `json:"product_view_count"`
	ProductPostedTime         string `json:"product_posted_time"`
	ProductPostedFriendlyTime int    `json:"product_posted_friendly_time"`
	ProductQuantity           int    `json:"product_quantity"`
	ProductQuantityUnits      string `json:"product_quantity_units"`
	ProductNotes              string `json:"product_notes"`
	PostedBy                  string `json:"posted_by"`
	ProductLatitude           string `json:"product_latitude"`
	ProductLongitude          string `json:"product_longitude"`
}

var items []Items

func allItems() {
	item := Items{
		ID:                        "1",
		ProductName:               "Desi Tomatoes",
		ProductDescription:        "Labelled as a vegetable for nutritional purposes, tomatoes are a good source of vitamin C and the phytochemical lycopene. The fruits are commonly eaten raw in salads, served as a cooked vegetable, used as an ingredient of various prepared dishes, and pickled.",
		ProductType:               "Food",
		ProductCategory:           "Groceries",
		ProductLikeCount:          15,
		ProductViewCount:          25,
		ProductPostedTime:         "02-09-2020 15:04",
		ProductPostedFriendlyTime: 20,
		ProductQuantity:           5,
		ProductQuantityUnits:      "Units",
		ProductNotes:              "Please collect it from watchman",
		PostedBy:                  "#12123",
		ProductLatitude:           "19.020054048637892",
		ProductLongitude:          "73.01573741893374",
	}
	item1 := Items{
		ID:                        "2",
		ProductName:               "Sony PlayStation 5",
		ProductDescription:        "Sony PS5 features ultra-high-speed SSD, adaptive triggers of the new DualSense controller, a slick new PS5 UI, and 3D Audio, and all series of incredible PS5",
		ProductType:               "Non-Food",
		ProductCategory:           "Gadgets",
		ProductLikeCount:          100,
		ProductViewCount:          200,
		ProductPostedTime:         "09-09-2020 17:04",
		ProductPostedFriendlyTime: 10,
		ProductQuantity:           1,
		ProductQuantityUnits:      "Units",
		ProductNotes:              "Please collect after 2pm",
		PostedBy:                  "#12123",
		ProductLatitude:           "19.020054048637892",
		ProductLongitude:          "73.01573741893374",
	}
	item2 := Items{
		ID:                        "3",
		ProductName:               "Electronic Guitar",
		ProductDescription:        "An electric guitar is a guitar that requires external amplification in order to be heard at typical performance volumes, unlike a standard acoustic guitar",
		ProductType:               "Non-Food",
		ProductCategory:           "Music",
		ProductLikeCount:          10,
		ProductViewCount:          20,
		ProductPostedTime:         "29-09-2020 18:27",
		ProductPostedFriendlyTime: 25,
		ProductQuantity:           2,
		ProductQuantityUnits:      "Units",
		ProductNotes:              "Please don't ring the bell",
		PostedBy:                  "#12123",
		ProductLatitude:           "19.020054048637892",
		ProductLongitude:          "73.01573741893374",
	}
	items = append(items, item, item1, item2)
	fmt.Println(items)
}
