package parent

import (
	"errors"
	"os"

	"github.com/mitchellh/go-ps"
)

const numberOfGenerationsToTry = 5

func Get() (Parent, error) {
	pid := os.Getpid()

	for range numberOfGenerationsToTry {
		process, err := ps.FindProcess(pid)

		if err != nil {
			return Parent{}, err
		}

		parentProcess, err := ps.FindProcess(process.PPid())

		if err != nil {
			return Parent{}, err
		}

		if term, err := match(parentProcess.Executable()); err == nil {
			return term, nil
		}

		pid = parentProcess.Pid()
	}

	return Parent{}, errors.New("terminal not found")
}
