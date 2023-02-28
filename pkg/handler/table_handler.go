package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	e "github.com/fpetrikovich/go-guestlist/pkg/exception"
	"github.com/fpetrikovich/go-guestlist/pkg/model"
	"github.com/fpetrikovich/go-guestlist/pkg/service"
)

type EventTableHandler struct {
	service service.IEventTableService
}

func NewEventTableHandler(ms service.IEventTableService) *EventTableHandler {
	return &EventTableHandler{service: ms}
}

/**
 * Create a table for the event.
 * CURL CMD: curl -X POST localhost:3000/tables -H 'Content-Type: application/json' -d '{ "capacity": 10 }'
 */
func (th *EventTableHandler) CreateTable(w http.ResponseWriter, r *http.Request) *e.AppError {
	var eTable model.EventTable

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	decoder := CreateBodyDecoder(r)
	err := decoder.Decode(&eTable)

	if err != nil {
		return e.ErrorCaseHanding(e.NewBadInputError("Unknown field in body"))
	}

	pTable, err := th.service.CreateTable(&eTable)

	if err != nil {
		return e.ErrorCaseHanding(err)
	}

	HandleJsonResponse(w, http.StatusOK, struct {
		Table    int `json:"id"`
		Capacity int `json:"capacity"`
	}{
		Table:    pTable.TableID,
		Capacity: pTable.Capacity,
	})

	return nil // success
}

/**
 * Fetch a table with the table id.
 * CURL CMD: curl -X GET localhost:3000/tables/{id}'
 */
func (th *EventTableHandler) GetTable(w http.ResponseWriter, r *http.Request) *e.AppError {

	params := mux.Vars(r)
	pId := params["id"]

	id, err := strconv.Atoi(pId)

	if err != nil {
		return &e.AppError{Error: err, Message: "[ERROR] Table ID is not a number.", Code: http.StatusBadRequest}
	}

	log.Print("[INFO] Fetching table with ID: ", id)

	eTable, err := th.service.GetTable(id)

	if err != nil {
		return &e.AppError{Error: err, Message: "[ERROR] Fetching table data.", Code: http.StatusBadRequest}
	}

	HandleJsonResponse(w, http.StatusOK, eTable)

	return nil // success
}

/**
 * Fetch all the tables of the event.
 * CURL CMD: curl -X GET localhost:3000/tables'
 */
func (th *EventTableHandler) GetTables(w http.ResponseWriter, r *http.Request) *e.AppError {

	log.Print("[INFO] Fetching tables...")

	retTables, err := th.service.GetTables()

	if err != nil {
		return &e.AppError{Error: err, Message: "[ERROR] Fetching tables unsuccessful.", Code: http.StatusInternalServerError}
	}

	HandleJsonResponse(w, http.StatusOK, struct {
		Tables []model.EventTable `json:"tables"`
	}{
		Tables: retTables,
	})

	return nil // success
}

/**
 * Get the sum of all the empty seats.
 * CURL CMD: curl -X GET localhost:3000/seats_empty'
 */
func (th *EventTableHandler) GetEmptySeats(w http.ResponseWriter, r *http.Request) *e.AppError {

	log.Print("[INFO] Empty seats count...")

	freeSeats, err := th.service.GetEmptySeats()

	if err != nil {
		return &e.AppError{Error: err, Message: "[ERROR] Fetching empty seat count.", Code: http.StatusInternalServerError}
	}

	HandleJsonResponse(w, http.StatusOK, struct {
		Seats int `json:"seats_empty"`
	}{
		Seats: freeSeats,
	})

	return nil // success
}
