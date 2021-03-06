package models

import (
	"fmt"
	"testing"
)

func TestNewTorrentSearch(t *testing.T) {
	InitDB(defaultConfig)

	s := NewTorrentSearch()
	s.Name("preacher")
	s.Season(2)
	s.Episode(2)

	r, err := s.Results(1)
	if err != nil {
		t.Error(err)
	}

	for _, e := range r.List {
		//fmt.Printf("%#v\n", e)
		fmt.Printf("%s %s %dx%d\n", e.Id.Hex(), *e.Name, *e.Season, *e.Episode)
	}
}
