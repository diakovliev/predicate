package predicate

// True returns a new predicate that always returns true
// regardless of the input.
func True[T any]() func(T) bool {
	return func(_ T) bool {
		return true
	}
}

// False returns a new predicate that always returns false
// regardless of the input.
func False[T any]() func(T) bool {
	return func(_ T) bool {
		return false
	}
}

// Const returns a new predicate that always returns the given value
// regardless of the input.
func Const[T any](v bool) func(T) bool {
	return func(_ T) bool {
		return v
	}
}

// Cond returns a new predicate that returns true if the given condition is true
// regardless of the input.
func Cond[T any](cond func() bool) func(T) bool {
	return func(_ T) bool {
		return cond()
	}
}

// Not returns a new predicate that is the logical NOT of the given predicate.
// The resulting predicate will return true if the given predicate returns false.
func Not[T any](predicate func(T) bool) func(T) bool {
	return func(t T) bool {
		return !predicate(t)
	}
}
