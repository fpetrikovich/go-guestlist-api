package exception

import "fmt"

type ExceedsCapacityError struct {
	Capacity  int
	ExceedsBy int
}

func (e *ExceedsCapacityError) Error() string {
	return fmt.Sprintf("Table has free capacity of %d, entourage exceeds by %d.", e.Capacity, e.ExceedsBy)
}

func NewExceedsCapacityError(capacity int, exceeds int) error {
	return &ExceedsCapacityError{
		Capacity:  capacity,
		ExceedsBy: exceeds,
	}
}
