package slices

type Predicate[T comparable] func(s []T) bool

func Contains[T comparable](test T) Predicate[T] {
	return func(s []T) (ret bool) {
		for _, v := range s {
			if v == test {
				ret = true
				break
			}
		}
		return
	}
}

func ContainsAny[T comparable](test ...T) Predicate[T] {
	return func(s []T) (ret bool) {
		for _, v := range test {
			if Contains(v)(s) {
				ret = true
				break
			}
		}
		return
	}
}

func ContainsAll[T comparable](test ...T) Predicate[T] {
	return func(s []T) (ret bool) {
		for _, v := range test {
			if !Contains(v)(s) {
				return
			}
		}
		ret = true
		return
	}
}

func HasPrefix[T comparable](test []T) Predicate[T] {
	return func(s []T) (ret bool) {
		if len(s) < len(test) {
			return
		}
		for i := 0; i < len(test); i++ {
			if s[i] != test[i] {
				return
			}
		}
		ret = true
		return
	}
}

func HasSuffix[T comparable](test []T) Predicate[T] {
	return func(s []T) (ret bool) {
		if len(s) < len(test) {
			return
		}
		for i := 0; i < len(test); i++ {
			if s[len(s)-len(test)+i] != test[i] {
				return
			}
		}
		ret = true
		return
	}
}
