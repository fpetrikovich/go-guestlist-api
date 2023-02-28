package service

import (
	"log"
	"time"

	e "github.com/fpetrikovich/go-guestlist/pkg/exception"
	"github.com/fpetrikovich/go-guestlist/pkg/model"
	"github.com/fpetrikovich/go-guestlist/pkg/repository"
)

/*
The methods of this service allow for the retrieval of guest data, creation of new guests,
updating existing guests, and deleting guests.

Additionally, this service checks if the number of accompanying guests is a valid input,
and checks if there is enough room at a table for the guests before creating or updating a guest.

The package also includes error handling for any exceptions that may occur during the process.
*/
type DefaultGuestService struct {
	guestRepository repository.IGuestRepository
	tableService    IEventTableService
}

func NewDefaultGuestService(gRepo repository.IGuestRepository, tService IEventTableService) *DefaultGuestService {
	return &DefaultGuestService{
		guestRepository: gRepo,
		tableService:    tService,
	}
}

func (d *DefaultGuestService) GetGuestList() ([]model.GuestData, error) {
	return d.guestRepository.GetGuestList()
}

func (d *DefaultGuestService) GetArrivedGuests() ([]model.GuestArrival, error) {
	return d.guestRepository.GetArrivedGuests()
}

func (d *DefaultGuestService) GetGuest(name string) (*model.Guest, error) {
	// Check name doesnt have spaces
	err := e.ValidateStringInput(name)
	if err != nil {
		return &model.Guest{}, err
	}
	return d.guestRepository.GetGuest(name)
}

/**
 * Creates a new guest to add to the guestlist. Checks if the guests fits at the specified
 * table, checking if the input parameters are valid.
 * If the guest and their entourage do not fit in the table, returns an ExceedsCapacity err.
 *
 * @param  params  pointer to GuestData
 */
func (d *DefaultGuestService) CreateGuest(params *model.GuestData) error {
	// Check entourage is a valid number
	err := e.ValidatePositiveInput(params.Accompanying_guests)
	if err != nil {
		return err
	}

	// Check name doesnt have spaces
	err = e.ValidateStringInput(params.Name)
	if err != nil {
		return err
	}

	// Check table capacity
	free, err := d.tableService.GetEmptySeatsAtTable(params.Table)
	if err != nil {
		return err
	}

	// No room at table
	if free < (params.Accompanying_guests + 1) {
		return e.NewExceedsCapacityError(free, (params.Accompanying_guests+1)-free)
	}

	return d.guestRepository.CreateGuest(params)
}

/**
 * Handle the arrival of a guest to the event.
 * Sets the guest as arrived if the new entourage still fits in the table. Sets the
 * guest as rejected if they no longer fit in the table. Updates the arrival time to now.
 *
 * @param  params  pointer to GuestData
 */
func (d *DefaultGuestService) UpdateGuest(params *model.GuestData) error {

	// Check entourage is a valid number
	err := e.ValidatePositiveInput(params.Accompanying_guests)
	if err != nil {
		return err
	}

	// fetch the guest to update
	guest, err := d.guestRepository.GetGuest(params.Name)
	if err != nil {
		return err
	}

	// difference between what was expected and who they brought ==> + if they brought more
	entourageDiff := params.Accompanying_guests - guest.Entourage
	freeSeats, err := d.guestRepository.GetGuestTableFreeSeats(params.Name)
	if err != nil {
		return err
	}

	// updating guest model
	guest.Entourage = params.Accompanying_guests
	guest.ArrivedAt = time.Now().Format("2006-01-02 15:04:05")

	// check if there is room for them in the table
	if freeSeats < entourageDiff {
		// update user to rejected
		guest.ArrivalStatus = model.Rejected
	} else {
		guest.ArrivalStatus = model.Arrived
	}

	log.Print("[INFO] Updating guest: ", *guest)

	return d.guestRepository.UpdateGuest(guest)
}

func (d *DefaultGuestService) DeleteGuest(name string) error {
	return d.guestRepository.DeleteGuest(name)
}
