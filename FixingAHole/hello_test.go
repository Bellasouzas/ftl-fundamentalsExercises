package hello_test

import (
	hello "github.com/bellasouzas/GoTDD/FixingAHole"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello world!"
	got := hello.Greeting()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}