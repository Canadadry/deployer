package ssh

import (
	"bytes"
	"path/filepath"
)

func (c *client) Run(path string, cmd string) (string, error) {
	session, err := c.ssh.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	b := &bytes.Buffer{}
	session.Stdout = b
	session.Stderr = b
	err = session.Run("cd " + filepath.Clean(path) + ";" + cmd)
	return b.String(), err
}
