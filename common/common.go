package common

import "github.com/diakovliev/go-predicates"

func Equal[T comparable](test T) predicates.Predicate[T] {
	return func(s T) bool {
		return s == test
	}
}
