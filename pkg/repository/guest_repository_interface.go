package repository

import "github.com/fpetrikovich/go-guestlist/pkg/model"

/*
This is an interface `IGuestRepository` for database logic regarding
*/
type IGuestRepository interface {
	// This method retrieves a list of all guests along with their data.
	GetGuestList() ([]model.GuestData, error)
	// This method retrieves a list of guests who have arrived at the event.
	GetArrivedGuests() ([]model.GuestArrival, error)
	// This method retrieves data of a single guest by their name.
	GetGuest(name string) (*model.Guest, error)
	// This method creates a new guest with the provided parameters.
	CreateGuest(params *model.GuestData) error
	// This method updates the data of a given guest.
	UpdateGuest(g *model.Guest) error
	// This method retrieves the number of free seats at a table assigned to a given guest.
	GetGuestTableFreeSeats(name string) (int, error)
	// This method deletes a guest by their name.
	DeleteGuest(name string) error
}
