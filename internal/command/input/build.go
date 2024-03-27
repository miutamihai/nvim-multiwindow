package input

import (
	"errors"
	"mihaimiuta/nvim-multiwindow/internal/terminal"
)

type CommandInput struct {
	Name string
	Args []string
}

func Build(term terminal.Terminal) (CommandInput, error) {
	switch term {
	case terminal.WezTerm:
		return CommandInput{Name: "/opt/homebrew/bin/wezterm-gui", Args: []string{"start"}}, nil
	default:
		return CommandInput{}, errors.New("unknown terminal passed")
	}
}
