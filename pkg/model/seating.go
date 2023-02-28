package model

/*
The `Seating` struct is a model in Go which represents a seating arrangement at an event. 

It has two fields:
- `TableID`: an integer that represents the table number.
- `GuestID`: an integer that represents the guest assigned to the table.

Both of these fields are used to represent the seating arrangement for a guest at a table in an event.
*/
type Seating struct {
	TableID int
	GuestID int
}
