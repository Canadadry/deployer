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

func (dr *dryrun) Run(path string, cmd string) (string, error) {
	out := fmt.Sprintf("runnning : %s in %s", cmd, path)
	fmt.Fprintln(dr.w, out)
	return out, nil
}

func (dr *dryrun) RunLocally(cmd string) (string, error) {
	out := fmt.Sprintf("runnning loccaly : %s", cmd)
	fmt.Fprintln(dr.w, out)
	return out, nil
}

func (dr *dryrun) Upload(from string, to string) error {
	fmt.Fprintf(dr.w, "upload from %s to %s\n", from, to)
	return nil
}

func (dr *dryrun) Download(from string, to string) error {
	fmt.Fprintf(dr.w, "download from %s to %s\n", from, to)
	return nil
}
