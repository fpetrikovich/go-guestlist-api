package service

import "github.com/fpetrikovich/go-guestlist/pkg/model"

/*
The `IGuestService` is an interface that defines methods for managing guest data.
It provides a way to abstract the implementation details of the guest service,
allowing different implementations to be swapped in and out as needed while still
adhering to the same interface.
*/
type IGuestService interface {
	// Retrieves a list of all guests represented by `[]model.GuestData`.
	GetGuestList() ([]model.GuestData, error)
	// Retrieves a list of arrived guests represented by `[]model.GuestArrival`.
	GetArrivedGuests() ([]model.GuestArrival, error)
	// Retrieves a single guest by name represented by a pointer to `model.Guest`.
	GetGuest(name string) (*model.Guest, error)
	// Creates a new guest with parameters represented by `model.GuestData`.
	CreateGuest(params *model.GuestData) error
	// Updates an existing guest with parameters represented by `model.GuestData`.
	UpdateGuest(params *model.GuestData) error
	// Deletes a guest by name.
	DeleteGuest(name string) error
}
