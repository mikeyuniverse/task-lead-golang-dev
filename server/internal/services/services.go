package services

import "grpc-server/internal/repo"

type Services struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Services {
	return &Services{
		repo: repo,
	}
}
