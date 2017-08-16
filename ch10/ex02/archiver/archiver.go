package archiver

import (
	"fmt"
	"io"
)

//
type format struct {
	name, magic string
	check       func(string) bool
	open        func(string, io.Writer) error
}

var formats []format

func RegisterFormat(name, magic string, check func(string) bool, open func(string, io.Writer) error) {
	formats = append(formats, format{name, magic, check, open})
}

func Open(src string, dst io.Writer) error {
	for _, f := range formats {
		if f.check(src) {
			return f.open(src, dst)
		}
	}
	return fmt.Errorf("undefine format: %s", src)
}
