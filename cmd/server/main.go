package main

import (
	"context"
	"grpc-practice/internal/server/config"
	"grpc-practice/internal/server/repo"
	"grpc-practice/internal/server/repo/mongo"
	server "grpc-practice/internal/server/server"
	"grpc-practice/internal/server/services"
	"log"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	mongo, err := mongo.Connect(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repo.New(mongo)
	if err != nil {
		log.Fatal(err)
	}

	services := services.New(repo)

	ctx := context.Background()
	srv, err := server.New(&ctx, config.Server, services)
	if err != nil {
		log.Fatal(err)
	}

	err = srv.Start()
	if err != nil {
		log.Fatal(err)
	}

}
