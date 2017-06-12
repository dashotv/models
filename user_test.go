package models

import "testing"

func TestCryptPassword(t *testing.T) {
	s := "blarg"
	e := "blarg"

	p := cryptPassword(s)
	if p != e {
		t.Errorf("cryptPassword %s == %s, want %s", s, p, e)
	}
}
