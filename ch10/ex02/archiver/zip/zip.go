package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/hrNakamura/go_learn/ch10/ex02/archiver"
)

const suffix = ".zip"
const magic = "PK\x03\x04"

func check(name string) bool {
	return strings.HasSuffix(strings.ToLower(name), suffix) || isZip(name)
}

func isZip(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()
	bMagic := []byte(magic)
	buf := make([]byte, 4)
	_, err = f.Read(buf)
	if err != nil || !bytes.Equal(buf, bMagic) {
		return false
	}
	return true
}

func open(src string, dst io.Writer) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		_, err = io.Copy(dst, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	archiver.RegisterFormat("zip", magic, check, open)
}
