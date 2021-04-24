package runner

import (
	"app/pkg/fs"
	"app/pkg/ssh"
	"bytes"
	"os/exec"
)

type Runner interface {
	Run(string, string) (string, error)
	RunLocally(string) (string, error)
	Upload(from string, to string) error
	Download(from string, to string) error
}

func New(s ssh.Ssh, l fs.FS) *runner {
	return &runner{
		ssh:   s,
		local: l,
	}
}

type runner struct {
	ssh   ssh.Ssh
	local fs.FS
}

func (r *runner) Run(path string, cmd string) (string, error) {
	out, err := r.ssh.Run(path, cmd)
	return out, err
}

func (r *runner) RunLocally(c string) (string, error) {
	buf := &bytes.Buffer{}
	cmd := exec.Command("bash", "-c", c)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	return buf.String(), err
}

func (r *runner) Upload(from string, to string) error {
	file, err := r.local.Open(from)
	if err != nil {
		return err
	}
	return r.ssh.Upload(to, file)
}

func (r *runner) Download(from string, to string) error {
	file, err := r.local.New(to)
	if err != nil {
		return err
	}
	return r.ssh.Download(from, file)
}
