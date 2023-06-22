# Go monadic

"Go monadic" is a functional programming library for Go. It's purpose is to provide
basic monad implementations to support functional programming approaches in Go.

## Monads

- Maybe

## API

This section describes the API of the library and how to use it.

### Maybe

The `Maybe` monad encapsulates a value and provides functions to `Map` the value and
to compose it with other functions returning a value of the `Maybe` type through `Bind`.
`Maybe` can be in two states: `Just` and `Nothing`. In the `Just` state it's containing
an actual value. In the `Nothing` state it's empty.

To create a value of `Maybe` there are two ways:

- `functional.Just[int, int](42)`: This creates a `Maybe` containing a value
- `functional.Nothing[int, int]`: This creates an empty `Maybe`
