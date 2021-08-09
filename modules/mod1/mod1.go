package mod1

import (
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}

	reply := "Hello " + name
	return reply, nil
}
