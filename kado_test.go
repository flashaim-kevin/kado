package kado

import (
	"github.com/flashaim-kevin/kado/src/config"
	"testing"
)

func TestNewKado(t *testing.T) {

	_, err := NewKado(config.SayInfo())

	if err != nil {
		t.Error("new fail")
	}
}

func TestKado_Hello(t *testing.T) {
	expected := "Hello world"

	kado, err := NewKado(config.SayInfo())
	if err != nil {
		t.Error("new fail")
	}

	got := kado.Hello()

	if expected != got {
		t.Errorf("Info error expected %s, but got %s", expected, got)
	}

}