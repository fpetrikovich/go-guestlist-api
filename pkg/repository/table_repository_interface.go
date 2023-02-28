package repository

import "github.com/fpetrikovich/go-guestlist/pkg/model"

/*
The `IEventTableRepository` interface defines a set of methods for managing event tables in an event management system.
*/
type IEventTableRepository interface {
	// Retrieves a list of all event tables.
	GetTables() ([]model.EventTable, error)
	// Retrieves the event table with the given id.
	GetTable(id int) (*model.EventTable, error)
	// Creates a new event table with the given parameters.
	CreateTable(table *model.EventTable) (*model.EventTable, error)
	// Deletes the event table with the given id.
	DeleteTable(id int) error
	// Retrieves the number of empty seats at a particular event table with the given id.
	GetEmptySeatsAtTable(id int) (int, error)
	// Retrieves the total number of empty seats across all event tables.
	GetEmptySeats() (int, error)
}
