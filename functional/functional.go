package functional

import "github.com/apfohl/go-monadic/maybe"

func Just[T, U any](value T) maybe.Maybe[T, U] {
	return maybe.Return[T, U](value)
}

func Nothing[T, U any]() maybe.Maybe[T, U] {
	return maybe.Nothing[T, U]{}
}
