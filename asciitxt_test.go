package asciitxt_test

import (
	"testing"

	"github.com/aldy505/asciitxt"
)

func TestNew(t *testing.T) {
	s := asciitxt.New("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG 0123456789 _ the quick brown fox jumps over the lazy dog")

	if s == "" {
		t.Error("should not be empty")
	}
}
func TestWithConfig(t *testing.T) {
	s := asciitxt.WithConfig("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG 0123456789 _ the quick brown fox jumps over the lazy dog", asciitxt.Config{})

	if s == "" {
		t.Error("should not be empty")
	}
}

func TestEmpty(t *testing.T) {
	s := asciitxt.New("")
	if s != "" {
		t.Error("should be empty, got:", s)
	}
}

func TestInvalidStyle(t *testing.T) {
	// this should panic
	assertPanic(t, func() { asciitxt.WithConfig("hello", asciitxt.Config{Style: 2}) })
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
