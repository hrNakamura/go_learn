package cycle

import "testing"

func TestCycle(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	var p interface{}
	p = &p
	a, b, c, d := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}, &link{value: "d"}
	a.tail, b.tail = b, a
	linkList := []link{*a, *b}
	unlinkList := []link{*c, *d}
	linkMap := make(map[int]link)
	linkMap[1] = *a
	linkMap[2] = *b
	unlinkMap := make(map[int]link)
	unlinkMap[1] = *c
	unlinkMap[2] = *d
	linkKeyMap := make(map[link]int)
	linkKeyMap[*a] = 1
	linkKeyMap[*b] = 2
	unlinkKeyMap := make(map[link]int)
	unlinkKeyMap[*c] = 1
	unlinkKeyMap[*d] = 2
	tests := []struct {
		x    interface{}
		want bool
	}{
		{1, false},
		{p, true},
		{a, true},
		{c, false},
		{linkList, true},
		{unlinkList, false},
		{linkMap, true},
		{unlinkMap, false},
		{linkKeyMap, true},
		{unlinkKeyMap, false},
	}
	for _, test := range tests {
		if Cycle(test.x) != test.want {
			t.Errorf("Cycle(%v) = %v", test.x, !test.want)
		}
	}
}
