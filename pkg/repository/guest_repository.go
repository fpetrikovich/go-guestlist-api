package repository

import (
	"database/sql"
	"fmt"
	"log"

	e "github.com/fpetrikovich/go-guestlist/pkg/exception"
	"github.com/fpetrikovich/go-guestlist/pkg/model"
)

/*
MySQL implementation of a guest repository.

The repository provides methods to retrieve, create, and update guest data in a MySQL database.

The repository struct `MySQLGuestRepository` contains a single field Connection which is a pointer to a `sql.DB` object
representing a connection to a MySQL database.

The methods also return any errors that may occur when interacting with the database and provide relevant logging output.
Errors are handled by the upper layers.
*/
type MySQLGuestRepository struct {
	Connection *sql.DB
}

func NewMySQLGuestRepository(connection *sql.DB) *MySQLGuestRepository {
	return &MySQLGuestRepository{
		Connection: connection,
	}
}

/**
 * Retrieves from the `guest` table all records, joining with the `seating` table
 * to get the table id. Returns an array of GuestData which includes the name,
 * entourage size, and table id.
 * Errors while scanning a row are notified, but not handled. This will mean only
 * rows that failed will have incomplete data instead of stopping all the operation.
 *
 * @return  array of GuestData
 */
func (db *MySQLGuestRepository) GetGuestList() ([]model.GuestData, error) {

	sqlStatement := `
		SELECT g.name, g.entourage, s.table_id
		FROM guest as g
		JOIN seating as s ON g.guest_id = s.guest_id;
	`
	rows, err := db.Connection.Query(sqlStatement)

	var guests []model.GuestData

	// Foreach guest
	for rows.Next() {
		var guest model.GuestData

		err = rows.Scan(&guest.Name, &guest.Accompanying_guests, &guest.Table)

		if err != nil {
			log.Print("[ERROR] ", err.Error())
		}

		guests = append(guests, guest)
	}
	return guests, nil
}

/**
 * Retrieves from the `guest` table all guest that have arrived at the event. Returns an
 * array of GuestArrival which includes the name, entourage size, and the arrival time.
 * Errors while scanning a row are notified, but not handled. This will mean only
 * rows that failed will have incomplete data instead of stopping all the operation.
 *
 * @return  array of GuestArrival
 */
func (db *MySQLGuestRepository) GetArrivedGuests() ([]model.GuestArrival, error) {

	sqlStatement := `
		SELECT name, entourage, arrived_at 
		FROM guest
		WHERE FIELD(arrival_status, "arrived", "left", "rejected")
	`
	rows, err := db.Connection.Query(sqlStatement)

	var guests []model.GuestArrival

	// Foreach guest
	for rows.Next() {
		var guest model.GuestArrival

		err = rows.Scan(&guest.Name, &guest.Accompanying_guests, &guest.Arrived_at)

		if err != nil {
			log.Print("[ERROR] ", err.Error())
		}

		guests = append(guests, guest)
	}
	return guests, nil
}

/**
 * Retrieves a guest from the `guest` table using the unique attribute name.
 * Returns said guest and handles the database error to return a custom exception.
 *
 * @param  name  name of the guest to fetch
 * @return       pointer to an instance of Guest
 */
func (db *MySQLGuestRepository) GetGuest(name string) (*model.Guest, error) {

	var guest model.Guest
	sqlStatement := `SELECT * FROM guest WHERE name = ?;`

	// Fetch record where the id matches
	row := db.Connection.QueryRow(sqlStatement, name)
	err := row.Scan(&guest.GuestID, &guest.Name, &guest.Entourage, &guest.ArrivalStatus, &guest.ArrivedAt, &guest.CreatedAt, &guest.UpdateAt)

	return &guest, e.CheckDatabaseError(err, name, "name", "guest")
}

/**
 * Inserts a new record in the `guest` table and uses the returned guest id to insert a record
 * in the `seating` table. Uses data from GuestData for the creation, which contains name, entourage
 * size, and table id. Returns nil if successful or a custom database exception upon an error.
 * If the guest already exists, a AlreadyExists error will occur.
 *
 * @param  params  pointer to GuestData
 */
func (db *MySQLGuestRepository) CreateGuest(params *model.GuestData) error {
	// insert the guest record into the mysql table
	res, err := db.Connection.Exec(`INSERT INTO guest (name, entourage) VALUES(?, ?);`, params.Name, params.Accompanying_guests)
	if err != nil {
		return e.CheckDatabaseError(err, params.Name, "name", "guest")
	}
	guestId, err := res.LastInsertId()
	if err != nil {
		return e.CheckDatabaseError(err, "", "", "")
	}

	// create the seating for the guest
	_, err = db.Connection.Exec(`INSERT INTO seating (guest_id, table_id) VALUES(?, ?);`, guestId, params.Table)
	if err != nil {
		return e.CheckDatabaseError(err, fmt.Sprint(guestId), "guestID", "guest")
	}

	return nil
}

/**
 * Updates a record in the `guest` table using the data from the instance of Guest.
 * If the guest id is not found, returns a NotFound error.
 *
 * @param  params  pointer to GuestData
 */
func (db *MySQLGuestRepository) UpdateGuest(guest *model.Guest) error {
	sqlStatement := `
		UPDATE guest
		SET
			name = ?,
			entourage = ?,
			arrival_status = ?,
			arrived_at = ?
		WHERE
			guest_id = ?
	`
	_, err := db.Connection.Exec(sqlStatement, guest.Name, guest.Entourage, guest.ArrivalStatus, guest.ArrivedAt, guest.GuestID)

	return e.CheckDatabaseError(err, fmt.Sprint(guest.GuestID), "guestID", "guest")
}

/**
 * Retrieves the free seats of the table the guest with name is sat at. Uses the `seating_usage`
 * view, retrieving the id of the table with a join of `guest` and `seating.`
 * Returns a NotFound error if name is not found.
 *
 * @param  name  name of the guest (string)
 * @return       free seats at the table where guest is
 */
func (db *MySQLGuestRepository) GetGuestTableFreeSeats(name string) (int, error) {

	var result int

	sqlStatement := `
		SELECT free_seats
		FROM seating_usage
		WHERE table_id = (
			SELECT s.table_id
			FROM guest as g
			JOIN seating as s ON g.guest_id = s.guest_id
			WHERE g.name = ?
		);
	`
	err := db.Connection.QueryRow(sqlStatement, name).Scan(&result)

	return result, e.CheckDatabaseError(err, name, "name", "guest")
}

/**
 * Given a guest name, deletes the guest (logically) by setting the arrival
 * status to 'left'. If guest is not found, returns NotFound. If no guest with
 * said name has arrived, returns ArrivalStatus error.
 *
 * @param  name  name of the guest (string)
 */
func (db *MySQLGuestRepository) DeleteGuest(name string) error {
	sqlStatement := `
		UPDATE guest
		SET arrival_status = 'left'
		WHERE name = ? AND FIELD(arrival_status, 'arrived');
	`
	res, err := db.Connection.Exec(sqlStatement, name)

	if err != nil {
		return e.CheckDatabaseError(err, name, "name", "guest")
	}

	n, err := res.RowsAffected()
	if err != nil {
		return e.CheckDatabaseError(err, name, "name", "guest")
	}

	if n == 0 {
		return e.NewArrivalStatusError("Guest can't leave before they arrive")
	}

	return err
}
