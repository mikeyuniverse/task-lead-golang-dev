package repo

import "grpc-server/internal/repo/mongo"

type Repo struct {
	db *mongo.MongoDB
}

func New(db *mongo.MongoDB) (*Repo, error) {
	return &Repo{db: db}, nil
}
