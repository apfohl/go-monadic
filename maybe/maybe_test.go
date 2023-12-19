package maybe_test

import (
	"github.com/apfohl/go-monadic/functional"
	"strconv"
	"testing"
)

func TestJust_create_maybe_with_value(t *testing.T) {
	input := 42
	expected := 42

	result := functional.Just(input).
		Match(
			func(value int) any {
				return value
			},
			func() any {
				t.FailNow()
				return 0
			},
		)

	if result != expected {
		t.Fatalf("Expected %d, but got %d", expected, result)
	}
}

func TestMap_maps_maybe_int_value_to_string(t *testing.T) {
	input := 42
	expected := "42"

	result := functional.Just(input).
		Map(func(i int) any {
			return strconv.Itoa(i)
		}).
		Match(
			func(value any) any {
				return value
			},
			func() any {
				t.FailNow()
				return 0
			},
		)

	if result != expected {
		t.Fatalf("Expected %s, but got %d", expected, result)
	}
}

func TestNothing_creates_empty_maybe(t *testing.T) {
	functional.Nothing[int]().
		Match(
			func(_ int) any {
				t.FailNow()
				return 0
			},
			func() any {
				return 42
			},
		)
}
