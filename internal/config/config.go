package config

import "os"

type Config struct {
	ServerPort    string
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func Load() Config {
	cfg := Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
	}
	return cfg
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
