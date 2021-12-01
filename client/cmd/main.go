package main

import (
	"grpc-client/internal/models"
	"grpc-client/internal/server"
	"log"
)

const url = "http://164.92.251.245:8080/api/v1/products"
const limit = 5     // List method request count
const itemLimit = 4 // Products count in one request

func main() {

	server := server.New()

	pagging := models.Pagging{
		Limit: limit,
	}

	sorting := models.Sorting{
		SortingType: models.SortingType{Alphabet: true},
		OrderType:   models.OrderType{Ascending: true},
	}

	// Download and update data in MongoDB
	server.Fetch(url)

	// Get pagging data
	for i := 0; i < limit; i++ {
		list, err := server.List(&pagging, &sorting)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(list)
	}
}
