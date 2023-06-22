package global

type FuncZero[TReturn any] func() TReturn

type Func[T, TReturn any] func(T) TReturn

type Action[TReturn any] func() TReturn

type Unit struct{}
