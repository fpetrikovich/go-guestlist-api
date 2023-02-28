package model

/*
The `EventTable` struct represents a model for a table in an event.

The struct has the following fields:
- `TableID`: an integer representing the ID of the table
- `Capacity`: an integer representing the maximum number of people that can sit at the table
- `CreatedAt`: a string representing the date and time when the table was created
- `UpdatedAt`: a string representing the date and time when the table was last updated

The json tags on each field are used for marshaling/unmarshaling the data to/from JSON,
so that when the data is encoded to JSON the keys in the JSON object will match the field names with the tags.
*/
type EventTable struct {
	TableID   int    `json:"id"`
	Capacity  int    `json:"capacity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"update_at"`
}
