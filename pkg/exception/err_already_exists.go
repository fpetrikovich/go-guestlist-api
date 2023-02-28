package exception

import "fmt"

type AlreadyExistsError struct {
	Id       string
	IdType   string
	Resource string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%s with %s %s already exists.", e.Resource, e.IdType, e.Id)
}

func NewAlreadyExistsError(id string, idType string, resource string) error {
	return &AlreadyExistsError{
		Id:       id,
		IdType:   idType,
		Resource: resource,
	}
}
