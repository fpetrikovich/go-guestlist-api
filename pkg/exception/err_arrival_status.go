package exception

import "fmt"

type ArrivalStatusError struct {
	Msg string
}

func (e *ArrivalStatusError) Error() string {
	return fmt.Sprintf("Invalid arrival status: %s", e.Msg)
}

func NewArrivalStatusError(msg string) error {
	return &ArrivalStatusError{
		Msg: msg,
	}
}
