package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

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

func ZipCheck(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("not string: %v", v)
	}
	if !regexp.MustCompile(`\d{5}`).Match([]byte(s)) {
		return fmt.Errorf("not zip code: %v", s)
	}
	return nil
}

func TestUnpack(t *testing.T) {
	type Unpacked struct {
		ZipCode string `http:"z" check:"zipC"` //`json:"v,omitempty"`
	}
	tests := []struct {
		req        *http.Request
		want       Unpacked
		errMessage string
	}{
		{
			&http.Request{Form: url.Values{"z": []string{"01234"}}},
			Unpacked{"01234"},
			"",
		},
		{
			&http.Request{Form: url.Values{"z": []string{"12A34"}}},
			Unpacked{""},
			"not zip code:",
		},
	}

	checks := map[string]URLCheck{
		"zipC": ZipCheck,
	}

	for _, test := range tests {
		var got Unpacked
		got.ZipCode = ""
		err := Unpack(test.req, &got, checks)
		if err != nil {
			if test.errMessage != "" && strings.Contains(err.Error(), test.errMessage) {
				t.Logf("Unpack(%+v) expect error: %v", test.req, err)
			} else {
				t.Errorf("Unpack(%+v) fail: %v", test.req, err)
			}
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Unpack(%+v): got %+v, want %+v", test.req, got, test.want)
		}
	}
}
