package mongoconnector

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connector struct {
	*mongo.Client
	*mongo.Database
	cfg Config
}

func New(cfg Config) *Connector {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(cfg.DB)

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	return &Connector{
		Client:   client,
		Database: database,
		cfg:      cfg,
	}
}
