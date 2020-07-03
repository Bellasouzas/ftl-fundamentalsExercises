package hello_test

import (
	hello "github.com/bellasouzas/GoTDD/PartOfThePackage"
	"testing"
)

func HelloTesting(t *testing.T) {
	got := hello.ReturnGreeting("Hi there")
	want := "Hi there yourself!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}