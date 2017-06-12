package models

import "testing"

func TestCryptPassword(t *testing.T) {
	s := "blarg"
	e := "01691bf2861daa645d88164a8479cdf6"

	p := cryptPassword(s)
	if p != e {
		t.Errorf("cryptPassword %s == %s, want %s", s, p, e)
	}
}
