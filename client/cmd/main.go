package main

import (
	"fmt"
	"grpc-client/internal/server"
	"log"
	"time"
)

const url = "http://164.92.251.245:8080/api/v1/products"

// const limit = 5     // List method request count
// const itemLimit = 4 // Products count in one request

func main() {

	server, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	// Download and update data in MongoDB
	for {
		time.Sleep(time.Second * 3)
		fmt.Println(server.Fetch(url))
	}

	// pagging := models.Pagging{
	// 	Limit: limit,
	// }

	// sorting := models.Sorting{
	// 	SortingType: models.SortingType{Alphabet: true},
	// 	OrderType:   models.OrderType{Ascending: true},
	// }

	// Get pagging data
	// for i := 0; i < limit; i++ {
	// for {
	// 	time.Sleep(time.Second * 5)
	// 	list, err := server.List(&pagging, &sorting)
	// 	if err != nil {
	// 		log.Println(err)
	// 		continue
	// 	}
	// 	log.Println(list)
	// }
}
