package gohello

import (
	"errors"
	"fmt"
)

func Crrr(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hello, %v. Welcome!", name)

	return message, nil
}
