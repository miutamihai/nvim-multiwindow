package terminal

import (
	"errors"
	"strings"
)

type Terminal string

const (
	WezTerm Terminal = "wezterm"
)

func match(input string) (Terminal, error) {
	if strings.Contains(input, string(WezTerm)) {
		return WezTerm, nil
	}

	return "", errors.New("No maching terminal found")
}
