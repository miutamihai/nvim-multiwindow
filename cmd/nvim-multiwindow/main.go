package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

type Terminal string

const (
	WezTerm Terminal = "wezterm"
)

func getRunningTerminal() (ps.Process, error) {
	pid := os.Getpid()

	for range 5 {
		process, err := ps.FindProcess(pid)

		if err != nil {
			return nil, err
		}

		parentProcess, err := ps.FindProcess(process.PPid())

		if err != nil {
			return nil, err
		}

		if strings.Contains(parentProcess.Executable(), string(WezTerm)) {
			return parentProcess, nil
		}

		pid = parentProcess.Pid()
	}

	return nil, errors.New("terminal not found")
}

func main() {
	terminal, err := getRunningTerminal()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Terminal: %s\n\twith pid:%d\n", terminal.Executable(), terminal.Pid())
}
