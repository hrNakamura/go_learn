// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestBool(t *testing.T) {
	tests := []struct {
		v    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for _, test := range tests {
		b, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		if string(b) != test.want {
			t.Errorf("Marshal(%v) get %s, want %s", test.v, b, test.want)
		}
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		v    float32
		want string
	}{
		{0, "0"},
		{1.5, "1.5"},
		{1.0, "1"},
		{1.5e+9, "1.5e+09"},
	}
	for _, test := range tests {
		b, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		if string(b) != test.want {
			t.Errorf("Marshal(%g) get %s, want %s", test.v, b, test.want)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		v    float64
		want string
	}{
		{0, "0"},
		{1.5, "1.5"},
		{1.0, "1"},
		{1.5e+9, "1.5e+09"},
	}
	for _, test := range tests {
		b, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		if string(b) != test.want {
			t.Errorf("Marshal(%g) get %s, want %s", test.v, b, test.want)
		}
	}
}

func TestComplex64(t *testing.T) {
	tests := []struct {
		v    complex64
		want string
	}{
		{1 + 1i, "#C(1 1)"},
		{1 - 1i, "#C(1 -1)"},
		{1, "#C(1 0)"},
	}
	for _, test := range tests {
		b, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		if string(b) != test.want {
			t.Errorf("Marshal(%g) get %s, want %s", test.v, b, test.want)
		}
	}
}

func TestComplex128(t *testing.T) {
	tests := []struct {
		v    complex128
		want string
	}{
		{1 + 1i, "#C(1 1)"},
		{1 - 1i, "#C(1 -1)"},
		{1, "#C(1 0)"},
	}
	for _, test := range tests {
		b, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		if string(b) != test.want {
			t.Errorf("Marshal(%g) get %s, want %s", test.v, b, test.want)
		}
	}
}

func TestInterface(t *testing.T) {
	type emptyStr struct {
		a, b, c interface{}
	}
	tests := []struct {
		v    emptyStr
		want string
	}{
		{emptyStr{1, 2, 3}, `((a ("interface {}" 1)) (b ("interface {}" 2)) (c ("interface {}" 3)))`},
	}
	for _, test := range tests {
		b, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		if string(b) != test.want {
			t.Errorf("Marshal(%v) get %s, want %s", test.v, b, test.want)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	type Values struct {
		B   bool
		F32 float32
		F64 float64
		I   interface{} `sexpr:"f"`
	}
	Interfaces["sexpr.Interface"] = reflect.TypeOf(int(0))
	tests := []struct {
		in   string
		want Values
	}{
		{`((B t) (F32 1.2) (F64 0) (I ("sexpr.Interface" 5)))`, Values{true, 1.2, 0, interface{}(5)}},
		{`((B nil) (F32 0) (F64 1.0) (f ("sexpr.Interface" 0)))`, Values{false, 0, 1, interface{}(0)}},
	}
	for _, test := range tests {
		var r Values
		err := Unmarshal([]byte(test.in), &r)
		if err != nil {
			t.Errorf("Unmarshal(%s) fail:%s", test.in, err)
		}
		if !reflect.DeepEqual(r, test.want) {
			t.Errorf("Unmarshal(%s) get %v, want %v", test.in, r, test.want)
		}
	}
}

func TestMarshal(t *testing.T) {
	type Values struct {
		B   bool    `sexpr:"b"`
		F32 float32 `sexpr:"f1"`
		F64 float64 `sexpr:"f2"`
		I   interface{}
	}
	tests := []struct {
		in   Values
		want string
	}{
		{Values{true, 1.5, 0, interface{}(5)}, `((b t) (f1 1.5) (f2 0) (I ("interface {}" 5)))`},
		{Values{false, 0, 1, interface{}(0)}, `((b nil) (f1 0) (f2 1) (I ("interface {}" 0)))`},
	}
	for _, test := range tests {
		data, err := Marshal(test.in)
		if err != nil {
			t.Errorf("Marshal(%+v) fail:%s", test.in, err)
		}
		got := string(data)
		if got != test.want {
			t.Errorf("Marshal(%+v) get %v, want %v", test.in, got, test.want)
		}
	}
}
