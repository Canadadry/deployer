package ssh

import (
	"github.com/pkg/sftp"
	"io"
)

func (c *client) Upload(pathTo string, from io.Reader) error {

	client, err := sftp.NewClient(c.ssh)
	if err != nil {
		return err
	}
	defer client.Close()

	to, err := client.Create(pathTo)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}

func (c *client) Download(pathFrom string, to io.Writer) error {

	client, err := sftp.NewClient(c.ssh)
	if err != nil {
		return err
	}
	defer client.Close()

	from, err := client.Open(pathFrom)
	if err != nil {
		return err
	}
	defer from.Close()

	_, err = io.Copy(to, from)
	return err
}
