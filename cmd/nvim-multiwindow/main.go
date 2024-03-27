package main

import (
	"errors"
	"mihaimiuta/nvim-multiwindow/internal/command/input"
	"mihaimiuta/nvim-multiwindow/internal/terminal"
	"os"
	"os/exec"
)

func main() {
	terminal, err := terminal.Get()

	if err != nil {
		panic(err)
	}

	commandInput, err := input.Build(terminal)

	if err != nil {
		panic(err)
	}

	if len(os.Args) != 2 {
		panic(errors.New("Needs exactly one argument"))
	}

	file := os.Args[1]

	args := append(commandInput.Args, "nvim", file, "-n")
	command := exec.Command(commandInput.Name, args...)
	_, startError := command.CombinedOutput()

	if startError != nil {
		panic(startError)
	}
}
