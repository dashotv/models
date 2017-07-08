package models

import (
	"fmt"

	"github.com/go-bongo/bongo"
)

const COLLECTION_USERS = "users"
const COLLECTION_MEDIA = "media"
const COLLECTION_TORRENTS = "torrents"
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
	torch, err := bongo.Connect(&bongo.Config{ConnectionString: host, Database: "torch_development"})
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", host, name, err))
	}
	media, err := bongo.Connect(&bongo.Config{ConnectionString: host, Database: "seer_development"})
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", host, name, err))
	}
	dashotv, err := bongo.Connect(&bongo.Config{ConnectionString: host, Database: "dashotv"})
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", host, name, err))
	}

	DB = &Connector{
		Users:    dashotv.Collection(COLLECTION_USERS),
		Media:    media.Collection(COLLECTION_MEDIA),
		Torrents: torch.Collection(COLLECTION_TORRENTS),
	}
}

func connect(db, host, collection string) (*bongo.Collection, error) {
	c, err := bongo.Connect(&bongo.Config{ConnectionString: host, Database: db})
	if err != nil {
		return nil, err
	}
	return c.Collection(collection), nil
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

type M map[string]interface{}
