package config

import "testing"

func TestSayInfo(t *testing.T) {
	expected := "Hello world"
	got := SayInfo()

	if expected != got {
		t.Errorf("Info error expected %s, but got %s", expected, got)
	}
}
