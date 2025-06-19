package config

import (
	"log"
	"os"
)

type Config struct {
	MongoURI  string
	DBName    string
	JWTSecret string
}

func Load() *Config {
	config := &Config{
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:    getEnv("DB_NAME", "dish_hub"),
		JWTSecret: getEnv("JWT_SECRET", "supersecretkey"),
	}

	logConfig(config)
	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func logConfig(config *Config) {
	log.Printf(
		"Configuration: MongoURI=%s, DBName=%s, JWTSecret=%s",
		config.MongoURI, config.DBName, config.JWTSecret,
	)
}
