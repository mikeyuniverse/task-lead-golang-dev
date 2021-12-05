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
		log.Println("Request \"Fetch\" with url - ", url)
		err = server.Fetch(url)
		if err != nil {
			log.Printf("FETCH ERROR: %s", err.Error())
		} else {
			log.Println("FETCH SUCCESS")
		}

		log.Println("Wait 5 seconds")
		time.Sleep(time.Second * 5)

		log.Println("Request \"List\"")
		list, err := server.List(&pagging, &sorting)
		if err != nil {
			log.Printf("Error - %s\n", err)
			continue
		}

		log.Println("Results - ", list)
		pagging.Start = pagging.Start + pagging.Limit
		log.Printf("New pagging params:\nStart - %d\nLimit - %d\n", pagging.Start, pagging.Limit)

		log.Println("Wait 3 seconds")
		time.Sleep(time.Second * 3)
	}
}
