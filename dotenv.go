package main

import (
	"github.com/joho/godotenv"
	"os"
)

type ENV struct {
	Port         string
	RootFilePath string
	DBUrl        string
}

func loadEnv() ENV {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	newEnv := ENV{}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	newEnv.Port = port

	rootFilePath := os.Getenv("ROOT_FILE_PATH")
	if rootFilePath == "" {
		rootFilePath = "./"
	}
	newEnv.RootFilePath = rootFilePath

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL env is not set")
	}
	newEnv.DBUrl = dbURL

	return newEnv

}
