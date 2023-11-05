package mongoconnector

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

type Connector struct {
	*mgo.Session
	*mgo.Database
	cfg Config
}

func New(cfg Config) *Connector {
	uri := fmt.Sprintf(
		"mongodb://%s:%s",
		cfg.Host,
		cfg.Port,
	)

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{uri},
		Database: "files",
		Username: cfg.Username,
		Password: cfg.Password,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	database := session.DB(cfg.DB)

	if err = session.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	return &Connector{
		Session:  session,
		Database: database,
		cfg:      cfg,
	}
}
