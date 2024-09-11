package strings

import (
	"regexp"
	"strings"
)

type Predicate func(s string) bool

func Contains(test string) Predicate {
	return func(s string) bool {
		return strings.Contains(s, test)
	}
}

func ContainsAny(test ...string) Predicate {
	return func(s string) bool {
		for _, v := range test {
			if Contains(v)(s) {
				return true
			}
		}
		return false
	}
}

func ContainsAll(test ...string) Predicate {
	return func(s string) bool {
		for _, v := range test {
			if !Contains(v)(s) {
				return false
			}
		}
		return true
	}
}

func HasPrefix(test string) Predicate {
	return func(s string) bool {
		return strings.HasPrefix(s, test)
	}
}

func HasSuffix(test string) Predicate {
	return func(s string) bool {
		return strings.HasSuffix(s, test)
	}
}

func Matches(test regexp.Regexp) Predicate {
	return func(s string) bool {
		return test.MatchString(s)
	}
}

func MatchesAny(test ...regexp.Regexp) Predicate {
	return func(s string) bool {
		for _, v := range test {
			if Matches(v)(s) {
				return true
			}
		}
		return false
	}
}

func MatchesAll(test ...regexp.Regexp) Predicate {
	return func(s string) bool {
		for _, v := range test {
			if !Matches(v)(s) {
				return false
			}
		}
		return true
	}
}
