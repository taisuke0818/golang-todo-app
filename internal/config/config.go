package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	MongoUri      string
	MongoDatabase string
	MongoUsername string
	MongoPassword string
}

func LoadConfig() (*Config, error) {
	var (
		out Config
		l   []string
	)
	if v, ok := os.LookupEnv("MONGO_URI"); !ok {
		l = append(l, "environment variable MONGO_URI is mandatory")
	} else {
		out.MongoUri = v
	}
	if v, ok := os.LookupEnv("MONGO_INITDB_DATABASE"); !ok {
		l = append(l, "environment variable MONGO_INITDB_DATABASE is mandatory")
	} else {
		out.MongoDatabase = v
	}
	if v, ok := os.LookupEnv("MONGO_INITDB_ROOT_USERNAME"); !ok {
		l = append(l, "environment variable MONGO_INITDB_ROOT_USERNAME is mandatory")
	} else {
		out.MongoUsername = v
	}
	if v, ok := os.LookupEnv("MONGO_INITDB_ROOT_PASSWORD"); !ok {
		l = append(l, "environment variable MONGO_INITDB_ROOT_PASSWORD is mandatory")
	} else {
		out.MongoPassword = v
	}
	if len(l) > 0 {
		return nil, fmt.Errorf("error: %s", strings.Join(l, ", "))
	}
	return &out, nil
}
