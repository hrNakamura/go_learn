package intset

type IntSet interface {
	Has(x int) bool
	Add(x int)
	UnionWith(t IntSet)
	String() string
	Ints() []int
}
