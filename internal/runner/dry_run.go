package runner

import (
	"fmt"
	"io"
)

type dryrun struct {
	w io.Writer
}

func NewDryRun(w io.Writer) *dryrun {
	return &dryrun{w: w}
}

func (dr *dryrun) Run(cmd string) error {
	fmt.Fprintf(dr.w, "runnning : %s\n", cmd)
	return nil
}

func (dr *dryrun) RunLocally(cmd string) error {
	fmt.Fprintf(dr.w, "runnning loccaly : %s\n", cmd)
	return nil
}

func (dr *dryrun) Upload(from string, to string) error {
	fmt.Fprintf(dr.w, "upload from %s to %s\n", from, to)
	return nil
}

func (dr *dryrun) Download(from string, to string) error {
	fmt.Fprintf(dr.w, "download from %s to %s\n", from, to)
	return nil
}
