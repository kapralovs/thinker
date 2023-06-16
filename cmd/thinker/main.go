package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	server "github.com/kapralovs/thinker/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	port := os.Getenv("API_PORT")
	app := server.NewApp()

	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
