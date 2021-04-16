package ssh

import (
	"github.com/pkg/sftp"
	"io"
)

func (c *client) Upload(path string, from io.Reader) error {

	client, err := sftp.NewClient(c.ssh)
	if err != nil {
		return err
	}
	defer client.Close()

	to, err := client.Create(path)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}

func (c *client) Download(path string, to io.Writer) error {

	client, err := sftp.NewClient(c.ssh)
	if err != nil {
		return err
	}
	defer client.Close()

	from, err := client.Open(path)
	if err != nil {
		return []byte{}, err
	}
	defer from.Close()

	_, err = io.Copy(to, from)
	return err
}
