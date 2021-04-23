package runner

import (
	"app/pkg/fs"
	"app/pkg/ssh"
	"io"
	"os"
	"os/exec"
)

type Runner interface {
	Run(string) error
	RunLocally(string) error
	Upload(from string, to string) error
	Download(from string, to string) error
}

func New() *runner {
	return &runner{
		Stdout:      os.Stdout,
		Stderr:      os.Stderr,
		Ssh:         nil,
		Local:       nil,
		DistantPath: "",
	}
}

type runner struct {
	Stdout      io.Writer
	Stderr      io.Writer
	Ssh         ssh.Ssh
	Local       fs.FS
	DistantPath string
}

func (r *runner) Run(cmd string) error {
	_, err := r.Ssh.Run(r.DistantPath, cmd)
	return err
}

func (r *runner) RunLocally(c string) error {
	cmd := exec.Command(c)
	cmd.Stdout = r.Stdout
	cmd.Stderr = r.Stderr
	return cmd.Run()
}

func (r *runner) Upload(from string, to string) error {
	file, err := r.Local.Open(from)
	if err != nil {
		return err
	}
	return r.Ssh.Upload(to, file)
}

func (r *runner) Download(from string, to string) error {
	file, err := r.Local.New(to)
	if err != nil {
		return err
	}
	return r.Ssh.Download(from, file)
}
