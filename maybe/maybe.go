package maybe

type Maybe[T any] interface {
	Match(just func(T) any, nothing func() any) any
	Map(mapper func(T) any) Maybe[any]
	Bind(mapper func(T) Maybe[any]) Maybe[any]
}

type Just[T any] struct {
	instance T
}

func Return[T any](value T) Maybe[T] {
	return Just[T]{instance: value}
}

func (maybe Just[T]) Match(just func(T) any, _ func() any) any {
	return just(maybe.instance)
}

func (maybe Just[T]) Map(mapper func(T) any) Maybe[any] {
	return Return[any](mapper(maybe.instance))
}

func (maybe Just[T]) Bind(mapper func(T) Maybe[any]) Maybe[any] {
	return mapper(maybe.instance)
}

type Nothing[T any] struct{}

func nothing[T any]() Maybe[T] {
	return Nothing[T]{}
}

func (maybe Nothing[T]) Match(_ func(T) any, nothing func() any) any {
	return nothing()
}

func (maybe Nothing[T]) Map(_ func(T) any) Maybe[any] {
	return nothing[any]()
}

func (maybe Nothing[T]) Bind(_ func(T) Maybe[any]) Maybe[any] {
	return nothing[any]()
}
