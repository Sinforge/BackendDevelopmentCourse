package main

import (
	"context"
	"log"
	"pr8/config"
	"pr8/internal/app"
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
