package predicates

// Predicate is a function that takes a type T and returns a boolean.
type Predicate[T any] func(T) bool

// True returns a new predicate that always returns true
// regardless of the input.
func True[T any]() Predicate[T] {
	return func(_ T) bool {
		return true
	}
}

// False returns a new predicate that always returns false
// regardless of the input.
func False[T any]() Predicate[T] {
	return func(_ T) bool {
		return false
	}
}

// Const returns a new predicate that always returns the given value
// regardless of the input.
func Const[T any](v bool) Predicate[T] {
	return func(_ T) bool {
		return v
	}
}

// Cond returns a new predicate that returns true if the given condition is true
// regardless of the input.
func Cond[T any](cond func() bool) Predicate[T] {
	return func(_ T) bool {
		return cond()
	}
}

// Not returns a new predicate that is the logical NOT of the given predicate.
// The resulting predicate will return true if the given predicate returns false.
func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !predicate(t)
	}
}
