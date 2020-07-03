package hello_test

import (
	hello "github.com/bellasouzas/GoTDD/epicfail"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Greeting()
	want := "Hello Gophers!"

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}

}