package tar

import (
	"archive/tar"
	"bytes"
	"io"
	"myProject/go_learn/ch10/ex02/archiver"
	"os"
	"strings"
)

const suffix = ".tar"
const magic = "ustar\x00"

func check(name string) bool {
	return strings.HasSuffix(strings.ToLower(name), suffix) || isTar(name)
}

func isTar(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()
	bMagic := []byte(magic)
	buf := make([]byte, 263)
	_, err = f.Read(buf)
	if err != nil || !bytes.Equal(buf[257:], bMagic) {
		return false
	}
	return true
}

func open(src string, dst io.Writer) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()
	tr := tar.NewReader(f)
	for {
		th, err := tr.Next()
		if err == io.EOF || th == nil {
			break
		} else if err != nil {
			return err
		}
		io.Copy(dst, tr)
	}
	return nil
}

func init() {
	archiver.RegisterFormat("tar", magic, check, open)
}
