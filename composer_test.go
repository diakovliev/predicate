package predicates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type row[T any] struct {
	i []bool
	o bool
}

func (r row[T]) I() (ret []Predicate[T]) {
	ret = make([]Predicate[T], len(r.i))
	for i, v := range r.i {
		ret[i] = Const[T](v)
	}
	return
}

var AND = []row[int]{
	{[]bool{true, true}, true},
	{[]bool{true, false}, false},
	{[]bool{false, true}, false},
	{[]bool{false, false}, false},
	// Edge cases...
	{[]bool{}, true},
	{[]bool{true}, true},
	{[]bool{false}, false},
}

var OR = []row[int]{
	{[]bool{true, true}, true},
	{[]bool{true, false}, true},
	{[]bool{false, true}, true},
	{[]bool{false, false}, false},
	// Edge cases...
	{[]bool{}, false},
	{[]bool{true}, true},
	{[]bool{false}, false},
}

var XOR = []row[int]{
	// 2 inputs...
	{[]bool{true, true}, false},
	{[]bool{true, false}, true},
	{[]bool{false, true}, true},
	{[]bool{false, false}, false},
	// 3 inputs...
	{[]bool{true, true, true}, true},
	{[]bool{true, true, false}, false},
	{[]bool{true, false, true}, false},
	{[]bool{true, false, false}, true},
	{[]bool{false, true, true}, false},
	{[]bool{false, true, false}, true},
	{[]bool{false, false, true}, true},
	{[]bool{false, false, false}, false},
	// Edge cases...
	{[]bool{}, false},
	{[]bool{true}, true},
	{[]bool{false}, false},
}

var ANYBUT = []row[int]{
	// 2 inputs...
	{[]bool{true, true}, false},
	{[]bool{true, false}, true},
	{[]bool{false, true}, true},
	{[]bool{false, false}, false},
	// 3 inputs...
	{[]bool{true, true, true}, false},
	{[]bool{true, true, false}, true},
	{[]bool{true, false, true}, true},
	{[]bool{true, false, false}, true},
	{[]bool{false, true, true}, true},
	{[]bool{false, true, false}, true},
	{[]bool{false, false, true}, true},
	{[]bool{false, false, false}, false},
	// Edge cases...
	{[]bool{}, false},
	{[]bool{true}, false},
	{[]bool{false}, false},
}

func TestComposers(t *testing.T) {
	type testCase struct {
		name     string
		composer Composer[int]
		cases    []row[int]
	}
	tests := []testCase{
		{
			name:     "And",
			composer: And[int],
			cases:    AND,
		},
		{
			name:     "Or",
			composer: Or[int],
			cases:    OR,
		},
		{
			name:     "Xor",
			composer: Xor[int],
			cases:    XOR,
		},
		{
			name:     "AnyBut",
			composer: AnyBut[int],
			cases:    ANYBUT,
		},
	}
	for _, tc := range tests {
		for _, r := range tc.cases {
			t.Run(fmt.Sprintf("%s(%v)==%t", tc.name, r.i, r.o), func(t *testing.T) {
				p := tc.composer(r.I()...)
				assert.Equal(t, r.o, p(0))
			})
		}
	}
}

func TestNot(t *testing.T) {
	assert.Equal(t, false, Not(func(int) bool { return true })(0))
	assert.Equal(t, true, Not(func(int) bool { return false })(0))
}
