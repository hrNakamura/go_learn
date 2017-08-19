package intset

import (
	"math/rand"
	"testing"
	"time"
)

func benchAdd(b *testing.B, size int, set IntSet) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			set.Add(j)
		}
	}
}

func benchUnionWith(b *testing.B, size int, rng *rand.Rand, setA, setB IntSet) {
	for i := 0; i < size; i++ {
		setA.Add(rng.Intn(2000))
		setB.Add(rng.Intn(2000))
	}
	for i := 0; i < b.N; i++ {
		setA.UnionWith(setB)
	}
}

func benchHas(b *testing.B, size int, set IntSet) {
	for i := 0; i < size; i++ {
		set.Add(i)
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			set.Has(j)
		}
	}
}

func benchString(b *testing.B, size int, set IntSet) {
	for i := 0; i < size; i++ {
		set.Add(i)
	}
	for i := 0; i < b.N; i++ {
		set.String()
	}
}

func BenchmarkWordAdd10(b *testing.B) {
	benchAdd(b, 10, &WordIntSet{})
}

func BenchmarkWordAdd100(b *testing.B) {
	benchAdd(b, 100, &WordIntSet{})
}

func BenchmarkWordAdd1000(b *testing.B) {
	benchAdd(b, 1000, &WordIntSet{})
}

func BenchmarkWordUnionWith10(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	benchUnionWith(b, 10, rng, &WordIntSet{}, &WordIntSet{})
}

func BenchmarkWordUnionWith100(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	benchUnionWith(b, 100, rng, &WordIntSet{}, &WordIntSet{})
}

func BenchmarkWordUnionWith1000(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	benchUnionWith(b, 1000, rng, &WordIntSet{}, &WordIntSet{})
}

func BenchmarkWordHas10(b *testing.B) {
	benchHas(b, 10, &WordIntSet{})
}

func BenchmarkWordHas100(b *testing.B) {
	benchHas(b, 100, &WordIntSet{})
}

func BenchmarkWordHas1000(b *testing.B) {
	benchHas(b, 1000, &WordIntSet{})
}

func BenchmarkWordString10(b *testing.B) {
	benchString(b, 10, &WordIntSet{})
}

func BenchmarkWordString100(b *testing.B) {
	benchString(b, 100, &WordIntSet{})
}

func BenchmarkWordString1000(b *testing.B) {
	benchString(b, 1000, &WordIntSet{})
}

func BenchmarkMapAdd10(b *testing.B) {
	benchAdd(b, 10, NewMapIntSet())
}

func BenchmarkMapAdd100(b *testing.B) {
	benchAdd(b, 100, NewMapIntSet())
}

func BenchmarkMapAdd1000(b *testing.B) {
	benchAdd(b, 1000, NewMapIntSet())
}

func BenchmarkMapUnionWith10(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	benchUnionWith(b, 10, rng, NewMapIntSet(), NewMapIntSet())
}

func BenchmarkMapUnionWith100(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	benchUnionWith(b, 100, rng, NewMapIntSet(), NewMapIntSet())
}

func BenchmarkMapUnionWith1000(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	benchUnionWith(b, 1000, rng, NewMapIntSet(), NewMapIntSet())
}

func BenchmarkMapHas10(b *testing.B) {
	benchHas(b, 10, NewMapIntSet())
}

func BenchmarkMapHas100(b *testing.B) {
	benchHas(b, 100, NewMapIntSet())
}

func BenchmarkMapHas1000(b *testing.B) {
	benchHas(b, 1000, NewMapIntSet())
}

func BenchmarkMapString10(b *testing.B) {
	benchString(b, 10, NewMapIntSet())
}

func BenchmarkMapString100(b *testing.B) {
	benchString(b, 100, NewMapIntSet())
}

func BenchmarkMapString1000(b *testing.B) {
	benchString(b, 1000, NewMapIntSet())
}

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
