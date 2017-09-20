package params

import "testing"

func TestPack(t *testing.T) {
	type unpacked struct {
		name string
		x    int
		y    bool
	}

	tests := []struct {
		s    unpacked
		want string
	}{
		{unpacked{"alpha", 1, true}, "name=alpha&x=1&y=true"},
		{unpacked{"beta", 0, false}, "name=beta&x=0&y=false"},
	}
	for _, test := range tests {
		u, err := Pack(&test.s)
		if err != nil {
			t.Errorf("Pack fail:%v", err)
		}
		if u.RawQuery != test.want {
			t.Errorf("Pack(%v): get %s, want %s", test.s, u.RawQuery, test.want)
		}
	}
}
