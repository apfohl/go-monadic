package maybe_test

import (
	"github.com/apfohl/go-monadic/functional"
	"testing"
)

func TestJust_create_maybe_with_value(t *testing.T) {
	input := 42
	expected := 42

	result := functional.Just[int, int](input).
		Match(
			func(value int) int {
				return value
			},
			func() int {
				t.FailNow()
				return 0
			},
		)

	if result != expected {
		t.Fatalf("Expected %d, but got %d", expected, result)
	}
}

func TestNothing_creates_empty_maybe(t *testing.T) {
	functional.Nothing[int, int]().
		Match(
			func(_ int) int {
				t.FailNow()
				return 0
			},
			func() int {
				return 42
			},
		)
}
