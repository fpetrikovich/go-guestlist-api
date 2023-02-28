package model

type GuestStatus string

// A constant string type that defines the possible arrival statuses of a guest.
const (
	NotArrived GuestStatus = "not_arrived"
	Arrived                = "arrived"
	Rejected               = "rejected"
	Left                   = "left"
	Allocate               = "allocate"
)

/*
The `Guest` struct is a model that represents a single guest at an event. 

It includes the following fields:
- `GuestID`: a unique identifier for the guest.
- `Name`: the name of the guest.
- `Entourage`: the number of guests accompanying the primary guest.
- `ArrivalStatus`: the status of the guest's arrival, represented as an instance of the GuestStatus type.
- `ArrivedAt`: the time when the guest arrived, stored as an interface type to accommodate different data types.
- `UpdateAt`: the time when the guest's information was last updated.
- `CreatedAt`: the time when the guest's information was created.

The json tags on each field are used for marshaling/unmarshaling the data to/from JSON, 
so that when the data is encoded to JSON the keys in the JSON object will match the field names with the tags.
*/
type Guest struct {
	GuestID       int         `json:"guest_id"`
	Name          string      `json:"name"`
	Entourage     int         `json:"accompanying_guest"`
	ArrivalStatus GuestStatus `json:"arrival_status"`
	ArrivedAt     interface{} `json:"arrived_at"`
	UpdateAt      string      `json:"updated_at"`
	CreatedAt     string      `json:"created_at"`
}

/*
The `GuestData` struct is a model representing data of a guest. 

It contains three fields:
- `Name`: A string representing the name of the guest.
- `Table`: An integer representing the table assigned to the guest.
- `Accompanying_guests`: An integer representing the number of guests accompanying the main guest.

The json tags on each field are used for marshaling/unmarshaling the data to/from JSON, 
so that when the data is encoded to JSON the keys in the JSON object will match the field names with the tags.
*/
type GuestData struct {
	Name                string `json:"name"`
	Table               int    `json:"table"`
	Accompanying_guests int    `json:"accompanying_guests"`
}

/*
The `GuestArrival` struct represents a model for storing information about a guest's arrival at an event. 

It contains the following fields:
- `Name`: A string representing the name of the guest.
- `Accompanying_guests`: An integer representing the number of guests accompanying the main guest.
- `ArrivedAt`: the time when the guest arrived, stored as an interface type to accommodate different data types.

The json tags on each field are used for marshaling/unmarshaling the data to/from JSON, 
so that when the data is encoded to JSON the keys in the JSON object will match the field names with the tags.
*/
type GuestArrival struct {
	Name                string `json:"name"`
	Accompanying_guests int    `json:"accompanying_guests"`
	Arrived_at          string `json:"time_arrived"`
}
