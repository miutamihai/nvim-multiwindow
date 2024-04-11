package input

import (
	"errors"
	"mihaimiuta/nvim-multiwindow/internal/parent"
)

type CommandInput struct {
	Name string
	Args []string
}

func Build(parentProcess parent.Parent) (CommandInput, error) {
	switch parentProcess.Program {
	case parent.WezTerm:
		return CommandInput{Name: "wezterm", Args: []string{"start"}}, nil
	case parent.Neovide:
		return CommandInput{Name: "neovide", Args: []string{"--", "-n"}}, nil
	default:
		return CommandInput{}, errors.New("unknown terminal passed")
	}
}
