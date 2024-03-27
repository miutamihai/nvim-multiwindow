package main

import (
	"fmt"
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

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	args := append(commandInput.Args, "nvim", cwd)

	pula := exec.Command(commandInput.Name, args...)
	output, startError := pula.CombinedOutput()
	fmt.Printf("%+v", string(output))

	if startError != nil {
		panic(startError)
	}
}
