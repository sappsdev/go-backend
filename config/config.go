package config

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
)

type UUID []byte

func Port() string {
	return getEnv("PORT", "4000")
}

func Host() string {
	return getEnv("HOST", "localhost")
}

func Debug() bool {
	value := os.Getenv("DEBUG")
	if len(value) == 0 {
		return true
	}
	return false
}

func MongoUri() string {
	uri := getEnv("MONGO_URI", "mongodb://localhost:27017")
	return fmt.Sprintf("%s", uri)
}

func Database() string {
	database := getEnv("DATABASE", "backend")
	return fmt.Sprintf("%s", database)
}

var SECRET string = "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"

func GenSecret() {
	SECRET = uuid.NewV4().String()
}

func FilesDir() string {
	debug := Debug()
	if(debug == false) {
		return fmt.Sprintf("%s", "./files")
	}
	return fmt.Sprintf("%s", "./dist/files")
}

func FilesAvatar() string {
	debug := Debug()
	if(debug == false) {
		return fmt.Sprintf("%s", "./files/avatar")
	}
	return fmt.Sprintf("%s", "./dist/files/avatar")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
