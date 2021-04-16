package ssh

import (
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"io"
	"net"
)

type Ssh interface {
	Run(path string, cmd string) (string, error)
	Upload(path string, file io.Reader) error
	Download(path string, file io.Writer) error
}

type Login struct {
	Addr                  string
	Port                  int
	User                  string
	PrivateKey            string
	Password              string
	InsecureIgnoreHostKey bool
}

type client struct {
	ssh *ssh.Client
}

func New(l Login) (*client, error) {

	config := &ssh.ClientConfig{
		User:            l.user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if l.InsecureIgnoreHostKey == false {
		hostKeyCallback, err := knownhosts.New(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
		if err != nil {
			return nil, err
		}
		config.HostKeyCallback = hostKeyCallback
	}
	if len(l.privateKey) == 0 && len(l.password) == 0 {
		return nil, fmt.Errorf("You must provite a password or private key")
	}
	if len(l.privateKey) > 0 {
		key, err := ssh.ParsePrivateKey([]byte(l.PrivateKey))
		if err != nil {
			return nil, err
		}

		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(key),
		}

	} else {
		config.Auth = []ssh.AuthMethod{
			ssh.Password(l.Password),
		}

	}

	c, err := ssh.Dial("tcp", net.JoinHostPort(l.addr, l.port), config)
	if err != nil {
		return nil, err
	}
	return &client{ssh: c}
}
