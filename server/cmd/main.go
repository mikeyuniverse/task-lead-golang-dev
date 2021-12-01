package main

import (
	"fmt"
	"grpc-server/internal/config"
	"grpc-server/internal/repo"
	"grpc-server/internal/repo/mongo"
	server "grpc-server/internal/server"
	"grpc-server/internal/services"
	"log"
)

func main() {

	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----CONFIGS")
	fmt.Printf("%+v\n", config)

	mongo, err := mongo.Connect(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----MONGO CONNECT")
	fmt.Printf("%+v\n", mongo)

	repo, err := repo.New(mongo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----REPO")
	fmt.Printf("%+v\n", repo)

	services := services.New(repo)

	fmt.Println("----SERVICES")
	fmt.Printf("%+v\n", services)

	srv, err := server.New(config.Server, services)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----SERVER")
	fmt.Printf("%+v\n", srv)

}
