package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type MongoConfig struct {
	Host           string
	Port           string
	DBName         string
	CollectionName string
}

func NewMongoConfig() (MongoConfig, error) {
	var mongo MongoConfig
	err := envconfig.Process("MONGO", &mongo)
	if err != nil {
		return MongoConfig{}, err
	}
	return mongo, nil
}

type GRPC struct {
	Port string
}

func NewGRPCConfig() (GRPC, error) {
	var grpc GRPC
	err := envconfig.Process("grpc", &grpc)
	if err != nil {
		return GRPC{}, err
	}
	return grpc, nil
}

type Config struct {
	DB     MongoConfig
	Server GRPC
}

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func Read() (*Config, error) {

	mongo, err := NewMongoConfig()
	if err != nil {
		return nil, err
	}

	grpc, err := NewGRPCConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB:     mongo,
		Server: grpc,
	}, nil
}
