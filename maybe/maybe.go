package maybe

import (
	"github.com/apfohl/go-monadic/global"
)

type Maybe[T, U any] interface {
	Match(just global.Func[T, U], nothing global.FuncZero[U]) U
	Map(mapper global.Func[T, U]) Maybe[U, T]
	Bind(mapper global.Func[T, Maybe[U, T]]) Maybe[U, T]
}

func Return[T, U any](value T) Maybe[T, U] {
	return Just[T, U]{instance: value}
}

func nothing[T, U any]() Maybe[T, U] {
	return Nothing[T, U]{}
}

type Just[T, _ any] struct {
	instance T
}

func (maybe Just[T, TResult]) Match(just global.Func[T, TResult], _ global.FuncZero[TResult]) TResult {
	return just(maybe.instance)
}

func (maybe Just[T, U]) Map(mapper global.Func[T, U]) Maybe[U, T] {
	return Return[U, T](mapper(maybe.instance))
}

func (maybe Just[T, U]) Bind(mapper global.Func[T, Maybe[U, T]]) Maybe[U, T] {
	return mapper(maybe.instance)
}

type Nothing[T, _ any] struct{}

func (maybe Nothing[T, TResult]) Match(_ global.Func[T, TResult], nothing global.FuncZero[TResult]) TResult {
	return nothing()
}

func (maybe Nothing[T, U]) Map(_ global.Func[T, U]) Maybe[U, T] {
	return nothing[U, T]()
}

func (maybe Nothing[T, U]) Bind(_ global.Func[T, Maybe[U, T]]) Maybe[U, T] {
	return nothing[U, T]()
}
