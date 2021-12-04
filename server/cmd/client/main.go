package main

import (
	"grpc-practice/internal/client/server"
	"grpc-practice/pkg/proto/transport"
	"log"
	"time"
)

const url = "http://164.92.251.245:8080/api/v1/products"
const grpcPort = "9000"

func main() {

	server, err := server.New(grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	pagging := transport.Pagging{
		Limit: 5,
		Start: 5,
	}

	sorting := transport.Sorting{
		SortingType: transport.SortingType{Name: true},
		OrderType:   transport.OrderType{Ascending: true},
	}

	for {
		time.Sleep(time.Second * 3)
		err = server.Fetch(url)
		if err != nil {
			log.Printf("FETCH ERROR: %s", err.Error())
		}
		log.Println("FETCH SUCCESS")

		time.Sleep(time.Second * 5)
		list, err := server.List(&pagging, &sorting)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(list)
		pagging.Start = pagging.Start + pagging.Limit
	}
}
