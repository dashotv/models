package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type TorrentResponse struct {
	BaseResponse `bson:",inline" json:",inline"`
	List         []*Torrent `json:"torrents"`
}

func (r *TorrentResponse) Add(t *Torrent) {
	r.List = append(r.List, t)
}

type Torrent struct {
	Document `bson:",inline"`

	Type   *string
	Source *string

	Raw         *string
	Title       *string
	Description *string
	Size        *string
	View        *string
	Download    *string
	Infohash    *string

	Name     *string
	Season   *int
	Episode  *int
	Volume   *int
	Checksum *string
	Group    *string
	Author   *string
	Verified bool

	Resolution *string
	Encoding   *string
	Quality    *string
	Widescreen bool
	Uncensored bool
	Bluray     bool

	Published time.Time `bson:"published_at"`
	Created   time.Time `bson:"created_at"`
	Updated   time.Time `bson:"updated_at"`
}

//type TorrentQuery struct {
//	BaseQuery
//	Exact bool
//}

//func (q *TorrentQuery) M() bson.M {
//	// TODO: handle non-exact queries
//	return q.Query
//}

func TorrentIndex(page int) (*TorrentResponse, error) {
	response := &TorrentResponse{}

	skip := (page - 1) * PER_PAGE

	results := DB.Torrents.Find(bson.M{})
	results.Query.Sort("-created_at")
	results.Query.Skip(skip)
	results.Query.Limit(PER_PAGE)

	for i := 0; i < PER_PAGE; i++ {
		t := &Torrent{}
		results.Next(t)
		response.Add(t)
	}

	pi, err := results.Paginate(PER_PAGE, page)
	if err != nil {
		return nil, err
	}
	response.Pagination(pi)

	return response, nil
}

//func NewTorrentQuery() *TorrentQuery {
//	return &TorrentQuery{
//		BaseQuery: BaseQuery{Query: bson.M{"verified": true}},
//		Exact:     true,
//	}
//}
//
//func TorrentFind(id string) (*Torrent, error) {
//	media := &Torrent{}
//
//	err := DB.Torrents.Find(id, media)
//	if err != nil {
//		return nil, err
//	}
//
//	return media, nil
//}
//
//func TorrentSearch(page int, query bson.M) (*TorrentResponse, error) {
//	r := &TorrentResponse{}
//
//	err := DB.Torrents.Where(query).Sort("-created_at").Page(page, r)
//	if err != nil {
//		return nil, err
//	}
//
//	return r, nil
//}
//
//func (m *Torrent) Save() error {
//	return DB.Torrents.Save(m)
//}
//
//func (m *Torrent) Find(id string) (*Torrent, error) {
//	err := DB.Torrents.Find(id, m)
//	if err != nil {
//		return nil, err
//	}
//
//	return m, nil
//}
