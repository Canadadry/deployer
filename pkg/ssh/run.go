package ssh

import (
	"path/filepath"
)

type byteBuf struct {
	content []byte
}

func (bb *byteBuf) Write(b []byte) (int, error) {
	bb.content = append(bb.content, b...)
	return len(b), nil
}

func (c *client) Run(path string, cmd string) (string, error) {
	session, err := c.ssh.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	b := &byteBuf{}
	session.Stdout = b
	session.Stderr = b
	err = session.Run("cd " + filepath.Clean(path) + ";" + cmd)
	return string(b.content), err
}
