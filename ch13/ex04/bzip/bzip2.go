package bzip

import (
	"io"
	"os/exec"
)

type Writer struct {
	cmd   exec.Cmd
	stdin io.WriteCloser
}

func NewWriter(w io.Writer) (io.WriteCloser, error) {
	cmd := exec.Cmd{Path: "/usr/bin/bzip2", Stdout: w}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	wc := &Writer{cmd, stdin}
	return wc, nil
}

func (w *Writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

func (w *Writer) Close() error {
	stdErr := w.stdin.Close()
	cmdErr := w.cmd.Wait()
	if stdErr != nil {
		return stdErr
	}
	if cmdErr != nil {
		return cmdErr
	}
	return nil
}
