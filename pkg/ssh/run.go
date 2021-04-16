package ssh

import (
	"bytes"
)

func (c *client) Run(path string, cmd string) (string, error) {
	session, err := c.ssh.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	b := &bytes.Buffer{}
	session.Stdout = b
	err = session.Run(cmd)
	return b.String(), err
}
