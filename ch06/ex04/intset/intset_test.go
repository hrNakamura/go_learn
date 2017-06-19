package intset

import "testing"

func TestElems(t *testing.T) {
	var s IntSet
	e := []int{1, 2, 3}
	s.AddAll(e...)

	if len(e) != len(s.Elems()) {
		t.Fatalf("Elems error %s", &s)
	}
	for i, v := range s.Elems() {
		if v != e[i] {
			t.Fatalf("Elems error %s", &s)
		}
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	var s IntSet
	s.AddAll(1, 2, 3, 4)

	var u IntSet
	u.AddAll(4, 8)

	s.SymmetricDifferenceWith(&u)
	if s.String() != "{1 2 3 8}" {
		t.Errorf("SymmetricDifferenceWith error %s, %s", &s, &u)
	}
	s.Clear()
	u.Clear()

	s.AddAll(4, 8)
	u.AddAll(1, 2, 3, 4)
	s.SymmetricDifferenceWith(&u)
	if s.String() != "{1 2 3 8}" {
		t.Errorf("SymmetricDifferenceWith error %s, %s", &s, &u)
	}
}

func TestDifferenceWith(t *testing.T) {
	var s IntSet
	s.AddAll(1, 2, 3, 4)

	var u IntSet
	u.AddAll(4, 8)

	s.DifferenceWith(&u)
	if s.String() != "{1 2 3}" {
		t.Errorf("DifferenceWith error %s, %s", &s, &u)
	}
	s.Clear()
	u.Clear()

	s.AddAll(4, 8)
	u.AddAll(1, 2, 3, 4)
	s.DifferenceWith(&u)
	if s.String() != "{8}" {
		t.Errorf("DifferenceWith error %s, %s", &s, &u)
	}
}

func TestIntersectWith(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)

	var u IntSet
	u.Add(2)

	s.IntersectWith(&u)
	if s.String() != u.String() {
		t.Errorf("IntersectWith Error %s, %s", &s, &u)
	}

	u.Add(3)
	s.IntersectWith(&u)
	if s.String() != "{2}" {
		t.Errorf("IntersectWith Error %s, %s", &s, &u)
	}
}

func TestAddAll(t *testing.T) {
	var s IntSet
	s.AddAll(1, 2, 3)
	if s.String() != "{1 2 3}" {
		t.Errorf("AddAll error %s", &s)
	}

	list := []int{4, 5, 6}
	s.AddAll(list...)
	if s.String() != "{1 2 3 4 5 6}" {
		t.Errorf("AddAll error %s", &s)
	}
}

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
		t.Errorf("%s not equal %s", &s, c)
	}
}
