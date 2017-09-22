package intset

import "testing"

//TODO　テストが体系だってない
//TODO　空実装でもテストが通過してしまう
//TODO　既存のテストと比較する
func TestUnionWith(t *testing.T) {
	intsets := []IntSet{&WordIntSet{}, NewMapIntSet()}
	var input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range input {
		intsets[i%2].Add(v)
	}

	intsets[0].UnionWith(intsets[1])
	intsets[1].UnionWith(intsets[0])

	if !compareInts(intsets[0].Ints(), input) || !compareInts(intsets[0].Ints(), intsets[1].Ints()) {
		t.Errorf("word %v, map %v, want %v", intsets[0].Ints(), intsets[1].Ints(), input)
	}
	if intsets[0].String() != intsets[1].String() {
		t.Errorf("word %s, map %s", intsets[0].String(), intsets[1].String())
	}
}

func TestOthers(t *testing.T) {
	intsets := []IntSet{&WordIntSet{}, NewMapIntSet()}
	var tests = []struct {
		input []int
		value int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{200}, 200},
		{[]int{}, 1},
	}
	for _, test := range tests {
		for _, in := range test.input {
			intsets[0].Add(in)
			intsets[1].Add(in)
		}
		if intsets[0].Has(test.value) != intsets[1].Has(test.value) {
			t.Errorf("word %v, map %v", intsets[0].Has(test.value), intsets[1].Has(test.value))
		}
		if intsets[0].String() != intsets[1].String() {
			t.Errorf("word %s, map %s", intsets[0].String(), intsets[1].String())
		}
	}
}

func compareInts(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}
