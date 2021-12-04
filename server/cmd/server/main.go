package main

import (
	"grpc-practice/internal/server/config"
	"grpc-practice/internal/server/repo"
	"grpc-practice/internal/server/repo/mongo"
	server "grpc-practice/internal/server/server"
	"grpc-practice/internal/server/services"
	"log"
)

func main() {

	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("----CONFIGS")
	// fmt.Printf("%+v\n", config)

	mongo, err := mongo.Connect(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("----MONGO CONNECT")
	// fmt.Printf("%+v\n", mongo)

	repo, err := repo.New(mongo)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("----REPO")
	// fmt.Printf("%+v\n", repo)

	services := services.New(repo)

	// fmt.Println("----SERVICES")
	// fmt.Printf("%+v\n", services)

	srv, err := server.New(config.Server, services)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("----SERVER")
	// fmt.Printf("%+v\n", srv)

	err = srv.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("SERVER STARTED")

}
