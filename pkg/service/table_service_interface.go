package service

import "github.com/fpetrikovich/go-guestlist/pkg/model"

/*
The `IEventTableService` is an interface that defines methods for managing event table data.
It provides a way to abstract the implementation details of the event table service,
allowing different implementations to be swapped in and out as needed while still
adhering to the same interface.
*/
type IEventTableService interface {
	// Retrieves a list of all event tables represented by `[]model.EventTable`.
	GetTables() ([]model.EventTable, error)
	// Retrieves a single event table by id represented by a pointer to `model.EventTable`.
	GetTable(id int) (*model.EventTable, error)
	// Creates a new event table with parameters represented by `model.EventTable`.
	CreateTable(table *model.EventTable) (*model.EventTable, error)
	// Deletes an event table by id.
	DeleteTable(id int) error
	// Retrieves the number of empty seats at a specific event table.
	GetEmptySeatsAtTable(id int) (int, error)
	// Retrieves the total number of empty seats across all event tables.
	GetEmptySeats() (int, error)
}
