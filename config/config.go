package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Mongo MongoDBConfig
}

type MongoDBConfig struct {
	MDB_ADDRESS string
	MDB_NAME    string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	return &Config{
		Mongo: MongoDBConfig{
			MDB_ADDRESS: cast.ToString(coalesce("MDB_ADDRESS", "mongodb://localhost:27017")),
			MDB_NAME:    cast.ToString(coalesce("MDB_NAME", "test")),
		},
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
