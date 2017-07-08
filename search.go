package models

import "strconv"

type Search struct {
}

func (s *Search) Int(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		i = 0
	}
	return i
}

func (s *Search) Bool(v string) bool {
	return "true" == v
}
