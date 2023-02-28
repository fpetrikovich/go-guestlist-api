package exception

import "fmt"

type NotFoundError struct {
	Id       string
	IdType   string
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with %s %s not found.", e.Resource, e.IdType, e.Id)
}

func NewNotFoundError(id string, idType string, resource string) error {
	return &NotFoundError{
		Id:       id,
		IdType:   idType,
		Resource: resource,
	}
}
