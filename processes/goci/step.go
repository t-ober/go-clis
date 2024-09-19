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

	// fmt.Println("(Regular) Executing", s.exe, s.args, "in", s.proj)

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	return s.message, nil
}
