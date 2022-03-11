package goexamples

import (
	"errors"
	"fmt"
)

func Greet(name string) (greeting string, err error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	_greeting := fmt.Sprintf("Good day %v, hope you're good", name)
	return _greeting, nil
}
