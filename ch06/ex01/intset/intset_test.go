package intset

import "testing"

func TestLen(t *testing.T) {
	length := 0
	var s IntSet
	if s.Len() != length {
		t.Errorf("%v length %d not equal %d\n", &s, s.Len(), length)
	}

	s.Add(1)
	length = 1
	if s.Len() != length {
		t.Errorf("%v length %d not equal %d\n", &s, s.Len(), length)
	}

	s.Add(1)
	length = 1
	if s.Len() != length {
		t.Errorf("%v length %d not equal %d\n", &s, s.Len(), length)
	}

	s.Add(2)
	length = 2
	if s.Len() != length {
		t.Errorf("%v length %d not equal %d\n", &s, s.Len(), length)
	}
}

func TestRemove(t *testing.T) {
	var s1 IntSet
	var s2 IntSet

	s1.Remove(0)
	if s1.String() != s2.String() {
		t.Errorf("%s not equal %s", &s1, &s2)
	}

	s1.Add(1)
	s1.Add(2)
	s2.Add(1)
	if s1.Remove(2); s1.String() != s2.String() {
		t.Errorf("%s not equal %s", &s1, &s2)
	}
}

func TestClear(t *testing.T) {
	var s IntSet
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("%s is not cleared", &s)
	}

	s.Add(1)
	s.Add(2)
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("%s is not cleared", &s)
	}
}

func TestCopy(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)

	//TODO
	c := s.Copy()
	if s.String() != c.String() {
		t.Errorf("%s not equal %s", &s, &c)
	}
}
