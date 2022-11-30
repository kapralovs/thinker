package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kapralovs/thinker/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	port := os.Getenv("API_PORT")
	app := server.NewApp()
	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
