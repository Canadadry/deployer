package runner

import ()

type Runner interface {
	Run(string) error
	RunLocally(string) error
	Upload(from string, to string) error
	Download(from string, to string) error
}

func New() *dryrun {
	return &dryrun{w: nil}
}
