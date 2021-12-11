package config

import (
	"os"
	"reflect"
	"testing"
)

type testTableSchema struct {
	name        string
	envFilename string
	envVars     map[string]string
	want        *Config
	wantErr     bool
}

func setEnv(env map[string]string) error {
	for key, value := range env {
		os.Setenv(key, value)
	}
	return nil
}

func TestRead(t *testing.T) {
	testTable := []testTableSchema{
		{
			name:        "OK",
			envFilename: "one.env",
			envVars: map[string]string{
				"MONGO_DBNAME":         "ПРАДУКТЫ",
				"MONGO_COLLECTIONNAME": "Коллекцио",
				"MONGO_HOST":           "127.0.0.0",
				"MONGO_PORT":           "27033",
				"GRPC_PORT":            "9",
			},
			want: &Config{
				DB: MongoConfig{
					Host:           "127.0.0.0",
					Port:           "27033",
					DBName:         "ПРАДУКТЫ",
					CollectionName: "Коллекцио",
				},
				Server: GRPC{
					Port: "9",
				},
			},
			wantErr: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			err := setEnv(test.envVars)
			if err != nil {
				t.Errorf("Error with setEnv. Error - %s", err.Error())
			}

			config, err := Read()
			if (err != nil) != test.wantErr {
				t.Errorf("\nWANT ERR. Read() error.\nWant - %v\nGot - %v\n", test.wantErr, err)
			}

			if !reflect.DeepEqual(config, test.want) {
				if !test.wantErr {
					t.Errorf("\nDEEP EQUAL. Read() error.\nWant - %v\nGot - %v\n", test.want, config)
				}
			}

		})
	}
}
