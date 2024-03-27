package terminal

import (
	"errors"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

func Get() (Terminal, error) {
	pid := os.Getpid()

	for range 5 {
		process, err := ps.FindProcess(pid)

		if err != nil {
			return "", err
		}

		parentProcess, err := ps.FindProcess(process.PPid())

		if err != nil {
			return "", err
		}

		if strings.Contains(parentProcess.Executable(), string(WezTerm)) {
			return WezTerm, nil
		}

		pid = parentProcess.Pid()
	}

	return "", errors.New("terminal not found")
}
