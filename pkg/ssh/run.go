package ssh

import (
	"bytes"
	"path/filepath"
	"sync"
)

type synchronizedWriter struct {
	m   sync.Mutex
	buf *bytes.Buffer
}

func (w *synchronizedWriter) Write(p []byte) (int, error) {
	w.m.Lock()
	defer w.m.Unlock()
	return w.buf.Write(p)
}

func (c *client) Run(path string, cmd string) (string, error) {
	session, err := c.ssh.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	b := &bytes.Buffer{}
	w := &synchronizedWriter{buf: b}
	session.Stdout = w
	session.Stderr = w
	err = session.Run("cd " + filepath.Clean(path) + ";" + cmd)
	return b.String(), err
}
