package main

import (
	"os"
)

//MongoConfig - Structure for mongo config parameters
type MongoConfig struct {
	host     string
	port     string
	database string
	username string
	password string
}

//MuxConfig - Structure for http server parameters
type MuxConfig struct {
	port string
}

var mongoConfig = MongoConfig{
	host:     getEnv("MONGO_HOST", "localhost"),
	port:     getEnv("MONGO_PORT", "27017"),
	database: getEnv("MONGO_DB", "airQuality"),
	username: getEnv("MONGO_USER", ""),
	password: getEnv("MONGO_PASSWORD", ""),
}

var muxConfig = MuxConfig{
	port: getEnv("SERVER_PORT", "12345"),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
