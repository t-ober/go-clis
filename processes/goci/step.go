package main

import (
	"os/exec"
)

type step struct {
	name    string
	exe     string
	args    []string
	message string
	proj    string
}

func newStep(name, exe, mes, proj string, args []string) step {
	return step{
		name:    name,
		exe:     exe,
		message: mes,
		proj:    proj,
		args:    args,
	}
}

func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  "go build",
			msg:   "go build failed",
			cause: err,
		}
	}

	return s.message, nil
}
