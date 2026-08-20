package config

import "os"

type Config struct {
	MongoURI     string
	Port         string
	DatabaseName string
}

func Load() Config {
	return Config{
		MongoURI:     getEnv("MONGO_URI", "mongodb://mongodb:27017"),
		Port:         getEnv("PORT", "8080"),
		DatabaseName: getEnv("DATABASE_NAME", "imgapp"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
