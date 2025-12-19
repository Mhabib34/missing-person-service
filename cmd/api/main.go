package main

import (
	"context"
	"log"
	"time"

	"github.com/Mhbib34/missing-person-service/cmd/wire"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env hanya sekali di awal
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	app, err := wire.InitializeServer()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	go app.Worker.Start(ctx, 5*time.Second)

	app.Router.Run(":3000")
}
