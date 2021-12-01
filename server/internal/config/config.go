package config

type MongoConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	DBName         string
	CollectionName string
}

func NewMongoConfig() (*MongoConfig, error) {
	return &MongoConfig{}, nil
}

type GRPC struct {
	Host string
	Port string
}

func NewGRPCConfig() (*GRPC, error) {
	return &GRPC{}, nil
}

type Config struct {
	DB     *MongoConfig
	Server *GRPC
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
