package functional

import "github.com/apfohl/go-monadic/maybe"

func Just[T any](value T) maybe.Maybe[T] {
	return maybe.Return[T](value)
}

func Nothing[T any]() maybe.Maybe[T] {
	return maybe.Nothing[T]{}
}
