package decode

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	tests := []struct {
		in         string
		want       []Token
		errMessage string
	}{
		{`(1 "a" (2))`, []Token{StartList{}, Int(1), String("a"), StartList{}, Int(2), EndList{}, EndList{}}, ""},
		{`(1 a)`, []Token{StartList{}, Int(1), Symbol("a"), EndList{}}, ""},
		{`(1 2`, []Token{StartList{}, Int(1), Int(2)}, "expect ')'"},
		{`1 a`, []Token{}, "expect '('"},
	}
	for _, test := range tests {
		decoder := NewDecoder(strings.NewReader(test.in))
		gets := make([]Token, 0)
		for {
			tok, err := decoder.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				if strings.Contains(err.Error(), test.errMessage) {
					t.Logf("Decode(%s) get expected error:%v, want:%s", test.in, err, test.errMessage)
					break
				} else {
					t.Errorf("decode(%s) fail: %s", test.in, err)
					break
				}
			}
			gets = append(gets, tok)
		}
		if !reflect.DeepEqual(gets, test.want) {
			t.Errorf("Decode(%s), got %s, want %s", test.in, gets, test.want)
		}
	}
}
