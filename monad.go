package monad

import "fmt"

type Maybe[T any] struct {
	value T
	just  bool
}

func (m Maybe[T]) String() string {
	if m.just {
		return fmt.Sprintf("Just(%v)", m.value)
	}

	return "Nothing"
}

func (m Maybe[T]) IsJust() bool {
	return m.just
}

func (m Maybe[T]) IsNothing() bool {
	return !m.just
}

func (m Maybe[T]) Unwrap() T {
	if m.IsNothing() {
		panic("try to unwrap a None value")
	}
	return m.value
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{
		value: value,
		just:  true,
	}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

// func Curry[T any, U any](f func(T, U, ...any), value T) func(U, ...any) {

// }

func Fmap[T any, U any](f func(T) U) func(Maybe[T]) Maybe[U] {
	return func(m Maybe[T]) Maybe[U] {
		if m.IsJust() {
			return Just(f(m.Unwrap()))
		}

		return Nothing[U]()
	}
}

func Join[T any](m Maybe[Maybe[T]]) Maybe[T] {
	if m.IsNothing() {
		return Nothing[T]()
	}

	return m.value
}

func Monad[T any, U any](m Maybe[T]) func(func(T) Maybe[U]) Maybe[U] {
	return func(f func(T) Maybe[U]) Maybe[U] {
		return Join(Fmap(f)(m))
	}
}
