package exception

import "fmt"

type BadInputError struct {
	Input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("Invalid input: %s", e.Input)
}

func NewBadInputError(input string) error {
	return &BadInputError{
		Input: input,
	}
}
