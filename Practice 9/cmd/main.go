package main

import (
	"context"
	"log"
	"pr9/config"
	"pr9/internal/app"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(config)

	if err = app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
