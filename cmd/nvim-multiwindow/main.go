package main

import (
	"errors"
	"mihaimiuta/nvim-multiwindow/internal/command/input"
	"mihaimiuta/nvim-multiwindow/internal/parent"
	"os"
	"os/exec"
)

func buildCommand(parent parent.Parent, file string) *exec.Cmd {
	commandInput, err := input.Build(parent)

	if err != nil {
		panic(err)
	}

	args := append(commandInput.Args, file)

	if parent.IsTerminal {
		args = append(commandInput.Args, "nvim", file, "-n")
	}

	return exec.Command(commandInput.Name, args...)
}

func main() {
	if len(os.Args) != 2 {
		panic(errors.New("Needs exactly one argument"))
	}

	file := os.Args[1]
	parent, err := parent.Get()

	if err != nil {
		panic(err)
	}

	_, startError := buildCommand(parent, file).CombinedOutput()

	if startError != nil {
		panic(startError)
	}
}
