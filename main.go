package models

import (
	"fmt"

	"github.com/go-bongo/bongo"
)

const COLLECTION_USERS = "Users"
const COLLECTION_MEDIA = "Media"
const COLLECTION_TORRENTS = "Torrents"
const PER_PAGE = 25

var (
	// DB holds the db connection
	DB *Connector
)

type Connector struct {
	connection *bongo.Connection
	Users      *bongo.Collection
	Media      *bongo.Collection
	Torrents   *bongo.Collection
}

type Document struct {
	bongo.DocumentBase `bson:",inline"`
}

func InitDB(name, host string) {
	config := &bongo.Config{
		ConnectionString: host,
		Database:         name,
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", host, name, err))
	}

	DB = &Connector{
		connection: connection,
		//Users:      NewCollector(COLLECTION_USERS, connection),
		//Media:      NewCollector(COLLECTION_MEDIA, connection),
		//Torrents:   NewCollector(COLLECTION_TORRENTS, connection),
		Users:    connection.Collection(COLLECTION_USERS),
		Media:    connection.Collection(COLLECTION_MEDIA),
		Torrents: connection.Collection(COLLECTION_TORRENTS),
	}
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}
