package models

type TorrentSearch struct {
	Search
	query M
	sort  string
	limit int
}

func NewTorrentSearch() *TorrentSearch {
	return &TorrentSearch{
		query: M{
			"verified": true,
		},
		sort:  "-published_at",
		limit: PER_PAGE,
	}
}

func (s *TorrentSearch) Sort(v string) {
	s.sort = v
}

func (s *TorrentSearch) Limit(v int) {
	s.limit = v
}

func (s *TorrentSearch) Page(p int) int {
	return (p - 1) * PER_PAGE
}

func (s *TorrentSearch) Name(v string) {
	if v != "" {
		s.query["name"] = v
	}
}

//func (s *TorrentSearch) NameContains(v string) {
//	s.query["name"] = regexp.MustCompile(v)
//}

func (s *TorrentSearch) Type(v string) {
	if v != "" {
		s.query["type"] = v
	}
}

func (s *TorrentSearch) Source(v string) {
	if v != "" {
		s.query["source"] = v
	}
}

func (s *TorrentSearch) Resolution(v string) {
	if v != "" {
		s.query["resolution"] = v
	}
}

func (s *TorrentSearch) Season(v int) {
	if v != 0 {
		s.query["season"] = v
	}
}

func (s *TorrentSearch) Episode(v int) {
	if v != 0 {
		s.query["episode"] = v
	}
}

func (s *TorrentSearch) Group(v string) {
	if v != "" {
		s.query["group"] = v
	}
}

func (s *TorrentSearch) Author(v string) {
	if v != "" {
		s.query["author"] = v
	}
}

func (s *TorrentSearch) Verified(v bool) {
	s.query["verified"] = v
}

func (s *TorrentSearch) Uncensored(v bool) {
	s.query["uncensored"] = v
}

func (s *TorrentSearch) Bluray(v bool) {
	s.query["bluray"] = v
}

func (s *TorrentSearch) Results(page int) (*TorrentResponse, error) {
	response := &TorrentResponse{}

	results := DB.Torrents.Find(s.query)
	results.Query.Sort(s.sort)
	results.Query.Skip(s.Page(page))
	results.Query.Limit(s.limit)

	for i := 0; i < s.limit; i++ {
		t := &Torrent{}
		if !results.Next(t) {
			break
		}
		response.Add(t)
	}

	pi, err := results.Paginate(s.limit, page)
	if err != nil {
		return nil, err
	}
	response.Pagination(pi)

	return response, nil
}
