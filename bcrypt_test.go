package gutil

import (
	"testing"
)

func TestCreatePassword(t *testing.T) {
	pw := "password"
	hash := CreatePassword([]byte(pw))

	if hash == "" {
		t.Error("not expecting a blank string")
	}
}
func TestComparePasswords(t *testing.T) {
	pw := "password"
	hash := CreatePassword([]byte(pw))
	if valid := ComparePasswords(hash, []byte(pw)); !valid {
		t.Errorf("got %t, want %t", valid, true)
	}
	if valid := ComparePasswords(hash, []byte("Password")); valid {
		t.Errorf("got %t, want %t", valid, false)
	}
	if valid := ComparePasswords(hash, []byte("")); valid {
		t.Errorf("got %t, want %t", valid, false)
	}
}
