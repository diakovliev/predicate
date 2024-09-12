package common

func Equal[T comparable](test T) func(s T) bool {
	return func(s T) bool {
		return s == test
	}
}
