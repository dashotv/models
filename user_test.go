package models

import "testing"

func TestUser_Save(t *testing.T) {
	u := &User{}
	u.Name = "Shawn"
	u.Email = "me@shawncatz.com"
	u.Password = "blarg"
	err := u.Save()
	if err != nil {
		t.Errorf("saving user: %s", err)
	}

	u2, err := UserFind("me@shawncatz.com")
	if err != nil {
		t.Errorf("error finding user: %s", err)
	}

	if !u2.CheckPassword("blarg") {
		t.Errorf("password doesn't match: %s != %s", cryptPassword("blarg"), u2.PasswordHash)
	}
}

func TestCryptPassword(t *testing.T) {
	s := "blarg"
	e := "01691bf2861daa645d88164a8479cdf6"

	p := cryptPassword(s)
	if p != e {
		t.Errorf("cryptPassword %s == %s, want %s", s, p, e)
	}
}
