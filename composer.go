package predicate

// Composer is a function that takes a variadic list of predicates and returns a single predicate.
// type Composer[T any] func(...func(T) bool) func(T) bool

// And returns a new predicate that is the logical AND of the given predicates.
// The resulting predicate will return true if all of the given predicates return true.
// The implementation guarantees what each predicate will be called no more than once.
// If one of the given predicates returns false, the resulting predicate will return false, and
// rest of the predicates will not be called.
//
// Edge cases:
//   - And() == True
//   - And(True) == True
//   - And(False) == False
func And[T any](predicates ...func(T) bool) func(T) bool {
	return func(t T) (and bool) {
		and = true
		for _, predicate := range predicates {
			if and = and && predicate(t); !and {
				break
			}
		}
		return
	}
}

// Or returns a new predicate that is the logical OR of the given predicates.
// The resulting predicate will return true if any of the given predicates return true.
// The implementation guarantees what each predicate will be called once.
//
// Edge cases:
//   - Or() == False
//   - Or(True) == True
//   - Or(False) == False
func Or[T any](predicates ...func(T) bool) func(T) bool {
	return func(t T) (or bool) {
		for _, predicate := range predicates {
			or = or || predicate(t)
		}
		return
	}
}

// Xor returns a new predicate that is the logical XOR of the given predicates.
// The implementation guarantees that each predicate will be called no more than once.
//
// Edge cases:
//   - Xor() == False
//   - Xor(True) == True
//   - Xor(False) == False
func Xor[T any](predicates ...func(T) bool) func(T) bool {
	return func(t T) bool {
		// In general XOR operation is equivelent to counting the number
		// of true inputs by mod 2. So...
		count := 0
		for _, predicate := range predicates {
			if predicate(t) {
				count++
			}
		}
		return count%2 == 1
	}
}

// AnyBut returns a new predicate that is the logical "any but not all" of the given predicates.
// In case of 2 predicates, this operation is equivalent to Xor, othervise, it is equivalent to
// And(Or(...), Not(And(...))).
// The implementation guarantees that each predicate will be called once.
//
// Edge cases:
//   - AnyBut() == False
//   - AnyBut(True) == False
//   - AnyBut(False) == False
func AnyBut[T any](predicates ...func(T) bool) func(T) bool {
	return func(t T) bool {
		and := true
		or := false
		for _, predicate := range predicates {
			v := predicate(t)
			and = and && v
			or = or || v
		}
		return or && !and
	}
}
