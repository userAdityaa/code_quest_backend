package config

import (
	"log"
	"os"
)

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func GetGithubClientID() string {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		log.Fatal("GITHUB_CLIENT_ID not found in .env")
	}
	return clientID
}

func GetGithubClientSecret() string {
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("GITHUB_CLIENT_SECRET not found in .env")
	}
	return clientSecret
}
