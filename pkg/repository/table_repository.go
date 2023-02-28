package repository

import (
	"database/sql"
	"fmt"
	"log"

	e "github.com/fpetrikovich/go-guestlist/pkg/exception"
	"github.com/fpetrikovich/go-guestlist/pkg/model"
)

/*
Provides a MySQL implementation of the `IEventTableRepository` interface.
The code allows for fetching and manipulating the `event_table` and `seating_usage`
tables in a MySQL database. All methods return an error variable for the upper
level to handle.
*/
type MySQLEventTableRepository struct {
	Connection *sql.DB
}

func NewMySQLEventTableRepository(connection *sql.DB) *MySQLEventTableRepository {
	return &MySQLEventTableRepository{
		Connection: connection,
	}
}

/**
 * Returns an array of model.EventTable registered in `event_table`.
 * If an error occurs while scanning a particular row, will log the error,
 * but continue scanning other rows.
 *
 * @return      array of event tables
 */
func (db *MySQLEventTableRepository) GetTables() ([]model.EventTable, error) {

	sqlStatement := `SELECT * FROM event_table;`
	rows, err := db.Connection.Query(sqlStatement)

	var tables []model.EventTable

	// Foreach table
	for rows.Next() {
		var eTable model.EventTable

		err = rows.Scan(&eTable.TableID, &eTable.Capacity, &eTable.CreatedAt, &eTable.UpdatedAt)

		if err != nil {
			log.Print("[ERROR] ", err.Error())
		}

		tables = append(tables, eTable)
	}
	return tables, nil
}

/**
 * Retrieves a record from `event_table` that matches the id passed in the parameters,
 * stores it in a model.EventTable instance, and returns the pointer to the instance.
 * If an error occurs, will compare the error to the possible databse error and return
 * the adecuate one.
 *
 * @param  id  id of the event table to fetch
 * @return     pointer to the instance of EventTable
 */
func (db *MySQLEventTableRepository) GetTable(id int) (*model.EventTable, error) {

	var eTable model.EventTable
	sqlStatement := `SELECT * FROM event_table WHERE table_id = ?;`

	// Fetch record where the id matches
	row := db.Connection.QueryRow(sqlStatement, id)
	err := row.Scan(&eTable.TableID, &eTable.Capacity, &eTable.CreatedAt, &eTable.UpdatedAt)

	return &eTable, e.CheckDatabaseError(err, fmt.Sprint(id), "tableID", "table")
}

/**
 * Given a pointer to an instance of EventTable, insert a record of it in the `event_table`
 * If properly added, the table id will be added to the instance. The pointer is returned.
 *
 * @param  table pointer to instance of EventTable with data to use in insertion
 * @return       pointer to instance of EventTable with TableID added
 */
func (db *MySQLEventTableRepository) CreateTable(table *model.EventTable) (*model.EventTable, error) {
	// insert the event table record into the mysql table
	res, err := db.Connection.Exec(`INSERT INTO event_table (capacity) VALUES(?);`, table.Capacity)
	if err != nil {
		return table, err
	}
	id, err := res.LastInsertId()

	if err != nil {
		return table, err
	}

	// update the model obj with the returned id before returning it
	table.TableID = int(id)

	return table, nil
}

/**
 * Return the remaining capacity at a table given the table id.
 * Uses the view seating_usage for the query.
 *
 * @param  id  id of the event table
 * @return     amount of free seats at the table
 */
func (db *MySQLEventTableRepository) GetEmptySeatsAtTable(id int) (int, error) {

	var result int

	sqlStatement := `
		SELECT free_seats
		FROM seating_usage
		WHERE table_id = ?;
	`
	err := db.Connection.QueryRow(sqlStatement, id).Scan(&result)

	return result, e.CheckDatabaseError(err, fmt.Sprint(id), "tableID", "table")
}

/**
 * Return the remaining capacity between all tables. Uses the `seating_usage` table.
 *
 * @return  amount of empty seats in all the tables
 */
func (db *MySQLEventTableRepository) GetEmptySeats() (int, error) {

	var result int

	sqlStatement := `SELECT SUM(free_seats) FROM seating_usage;`

	err := db.Connection.QueryRow(sqlStatement).Scan(&result)

	return result, e.CheckDatabaseError(err, "", "", "free seats")
}

/**
 * To implement in the future.
 * Should delete the table from the database. If table had any guests, their arrival status
 * must be updated to allocate (since they need to be allocated to a new table).
 *
 * @param  id  id of the event table to delete
 */
func (db *MySQLEventTableRepository) DeleteTable(id int) error {
	panic("Implement me!")
}
