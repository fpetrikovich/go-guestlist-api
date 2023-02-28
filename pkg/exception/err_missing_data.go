package exception

import "fmt"

type MissingDataError struct {
	DataType string
}

func (e *MissingDataError) Error() string {
	return fmt.Sprintf("No data to show when querying for %s.", e.DataType)
}

func NewMissingDataError(dataType string) error {
	return &MissingDataError{
		DataType: dataType,
	}
}
