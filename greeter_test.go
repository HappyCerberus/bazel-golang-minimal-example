package greeter

import (
	"testing"
)

func TestGreeter(t *testing.T) {
	got := Greet()
	want := "Hello, whoever!"
	if got != want {
		t.Errorf(`Greet() = %q, want %q`, got, want)
	}
}
