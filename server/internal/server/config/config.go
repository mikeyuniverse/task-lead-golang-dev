package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type MongoConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	DBName         string
	CollectionName string
}

func NewMongoConfig() (*MongoConfig, error) {
	var mongo MongoConfig
	err := envconfig.Process("MONGO", &mongo)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", mongo)
	return &mongo, nil
}

type GRPC struct {
	Host string
	Port string
}

func NewGRPCConfig() (*GRPC, error) {
	var grpc GRPC
	err := envconfig.Process("grpc", &grpc)
	if err != nil {
		return nil, err
	}
	return &grpc, nil
}

type Config struct {
	DB     *MongoConfig
	Server *GRPC
}

func Read() (*Config, error) {
	err := godotenv.Load(".env")
	time.Sleep(time.Millisecond * 500)
	if err != nil {
		return nil, err
	}
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
