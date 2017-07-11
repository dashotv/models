package models

import (
	"fmt"

	"github.com/go-bongo/bongo"
)

const PER_PAGE = 25

var (
	// DB holds the db connection
	DB            *Connector
	defaultConfig = &Config{
		Host: "127.0.0.1",
		Torrents: &ConfigEntry{
			Database:   "torch_development",
			Collection: "torrents",
		},
		Media: &ConfigEntry{
			Database:   "seer_development",
			Collection: "media",
		},
		Users: &ConfigEntry{
			Database:   "dashotv",
			Collection: "users",
		},
	}
)

type Config struct {
	Host     string
	Torrents *ConfigEntry
	Media    *ConfigEntry
	Users    *ConfigEntry
}

type ConfigEntry struct {
	Database   string
	Collection string
}

type Connector struct {
	config     *Config
	connection *bongo.Connection
	Users      *bongo.Collection
	Media      *bongo.Collection
	Torrents   *bongo.Collection
}

type Document struct {
	bongo.DocumentBase `bson:",inline"`
}

func InitDB(c *Config) {
	torch, err := bongo.Connect(&bongo.Config{ConnectionString: c.Host, Database: c.Torrents.Database})
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", c.Host, c.Torrents.Database, err))
	}
	media, err := bongo.Connect(&bongo.Config{ConnectionString: c.Host, Database: c.Media.Database})
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", c.Host, c.Media.Database, err))
	}
	dashotv, err := bongo.Connect(&bongo.Config{ConnectionString: c.Host, Database: c.Users.Database})
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", c.Host, c.Users.Database, err))
	}

	DB = &Connector{
		config:   c,
		Users:    dashotv.Collection(c.Users.Collection),
		Media:    media.Collection(c.Media.Collection),
		Torrents: torch.Collection(c.Torrents.Collection),
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
