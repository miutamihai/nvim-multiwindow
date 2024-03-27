package terminal

import (
	"errors"
	"os"

	"github.com/mitchellh/go-ps"
)

const numberOfGenerationsToTry = 5

func Get() (Terminal, error) {
	pid := os.Getpid()

	for range numberOfGenerationsToTry {
		process, err := ps.FindProcess(pid)

		if err != nil {
			return "", err
		}

		parentProcess, err := ps.FindProcess(process.PPid())

		if err != nil {
			return "", err
		}

		if term, err := match(parentProcess.Executable()); err == nil {
			return term, nil
		}

		pid = parentProcess.Pid()
	}

	return "", errors.New("terminal not found")
}
