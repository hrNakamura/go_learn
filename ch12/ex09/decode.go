package decode

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token interface{}
type Symbol string
type String string
type Int int
type StartList struct{}
type EndList struct{}

type Decoder struct {
	scan  scanner.Scanner
	depth int
}

func NewDecoder(r io.Reader) *Decoder {
	var scan scanner.Scanner
	scan.Init(r)
	decoder := Decoder{scan: scan}
	return &decoder
}

func (d *Decoder) Token() (Token, error) {
	t := d.scan.Scan()
	if d.depth == 0 && t != '(' && t != scanner.EOF {
		return nil, fmt.Errorf("expect '(', get %s", scanner.TokenString(t))
	}
	switch t {
	case scanner.EOF:
		if d.depth != 0 {
			return nil, fmt.Errorf("expect ')', get %s", scanner.TokenString(t))
		}
		return nil, io.EOF
	case scanner.Ident:
		return Symbol(d.scan.TokenText()), nil
	case scanner.String:
		text := d.scan.TokenText()
		return String(text[1 : len(text)-1]), nil
	case scanner.Int:
		n, err := strconv.ParseInt(d.scan.TokenText(), 10, 64)
		if err != nil {
			return nil, err
		}
		return Int(n), nil
	case '(':
		d.depth++
		return StartList{}, nil
	case ')':
		d.depth--
		return EndList{}, nil
	default:
		return nil, fmt.Errorf("unexpected token %s at %v", scanner.TokenString(t), d.scan.Pos())
	}
}
