package parent

import (
	"errors"
	"strings"
)

type Parent struct {
	Program    Program
	IsTerminal bool
}

type Program string

const (
	WezTerm Program = "wezterm"
	Neovide Program = "neovide"
)

func match(input string) (Parent, error) {
	if strings.Contains(input, string(WezTerm)) {
		return Parent{
			Program:    WezTerm,
			IsTerminal: true,
		}, nil
	}

	if strings.Contains(input, string(Neovide)) {
		return Parent{
			Program:    Neovide,
			IsTerminal: false,
		}, nil
	}

	return Parent{}, errors.New("No maching parent found")
}
