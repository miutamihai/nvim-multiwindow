package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

type Terminal string

const (
	WezTerm Terminal = "wezterm"
)

func getRunningTerminal() (Terminal, error) {
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

type CommandInput struct {
	name string
	args []string
}

func getCommandInput(terminal Terminal) (CommandInput, error) {
	switch terminal {
	case WezTerm:
		return CommandInput{name: "/opt/homebrew/bin/wezterm-gui", args: []string{"start"}}, nil
	default:
		return CommandInput{}, errors.New("unknown terminal passed")
	}
}

func main() {
	terminal, err := getRunningTerminal()

	if err != nil {
		panic(err)
	}

	commandInput, err := getCommandInput(terminal)

	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	args := append(commandInput.args, "nvim", cwd)

	pula := exec.Command(commandInput.name, args...)
	output, startError := pula.CombinedOutput()
	fmt.Printf("%+v", string(output))

	if startError != nil {
		panic(startError)
	}
}
