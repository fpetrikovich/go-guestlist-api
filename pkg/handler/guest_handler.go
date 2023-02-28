package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	e "github.com/fpetrikovich/go-guestlist/pkg/exception"
	"github.com/fpetrikovich/go-guestlist/pkg/model"
	"github.com/fpetrikovich/go-guestlist/pkg/service"
)

type GuestHandler struct {
	service service.IGuestService
}

func NewGuestHandler(ms service.IGuestService) *GuestHandler {
	return &GuestHandler{service: ms}
}

/**
 * Retrieve the guest with name {name}
 * CURL EX: curl -X GET localhost:3000/guest_list/{name}'
 */
func (gh *GuestHandler) GetGuest(w http.ResponseWriter, r *http.Request) *e.AppError {

	// retrieve path params
	params := mux.Vars(r)
	name := params["name"]

	log.Print("[INFO] Fetching guest with name: ", name)

	guest, err := gh.service.GetGuest(name)

	if err != nil {
		return e.ErrorCaseHanding(err)
	}

	HandleJsonResponse(w, http.StatusOK, guest)

	return nil // success
}

/**
 * Retrieve all the guests
 * CURL EX: curl -X GET localhost:3000/guest_list'
 */
func (gh *GuestHandler) GetGuestList(w http.ResponseWriter, r *http.Request) *e.AppError {

	log.Print("[INFO] Fetching guest list...")

	guests, err := gh.service.GetGuestList()

	if err != nil {
		return e.ErrorCaseHanding(err)
	}

	HandleJsonResponse(w, http.StatusOK, struct {
		Guests []model.GuestData `json:"guests"`
	}{
		Guests: guests,
	})

	return nil // success
}

/**
 * Retrieve the list of all the arrived guests.
 * CURL CMD: curl -X GET localhost:3000/guests'
 */
func (gh *GuestHandler) GetArrivedGuests(w http.ResponseWriter, r *http.Request) *e.AppError {

	log.Print("[INFO] Fetching arrived guests...")

	guests, err := gh.service.GetArrivedGuests()

	if err != nil {
		return e.ErrorCaseHanding(err)
	}

	HandleJsonResponse(w, http.StatusOK, struct {
		Guests []model.GuestArrival `json:"guests"`
	}{
		Guests: guests,
	})

	return nil // success
}

/**
 * Adds a guest to the guest list.
 * CURL CMD: curl -X POST localhost:3000/guest_list -H 'Content-Type: application/json' -d '{"table": int, "accompanying_guests": int}'
 */
func (gh *GuestHandler) CreateGuest(w http.ResponseWriter, r *http.Request) *e.AppError {

	var bodyParams model.GuestData

	pathParams := mux.Vars(r)
	bodyParams.Name = pathParams["name"]

	log.Printf("[INFO] Creating guest %s...", bodyParams.Name)

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	decoder := CreateBodyDecoder(r)
	err := decoder.Decode(&bodyParams)

	if err != nil {
		return e.ErrorCaseHanding(e.NewBadInputError("Unknown field in body"))
	}

	err = gh.service.CreateGuest(&bodyParams)

	if err != nil {
		return e.ErrorCaseHanding(err)
	}

	HandleJsonResponse(w, http.StatusOK, struct {
		Name string `json:"name"`
	}{Name: bodyParams.Name})

	return nil // success
}

/**
 * Set a guest as arrived.
 * CURL CMD: curl -X PUT "localhost:3000/guests/<name>" -H 'Content-Type: application/json' -d '{"accompanying_guests": int}'
 */
func (gh *GuestHandler) UpdateGuest(w http.ResponseWriter, r *http.Request) *e.AppError {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	var bodyParams model.GuestData
	bodyParams.Name = name

	decoder := CreateBodyDecoder(r)
	err := decoder.Decode(&bodyParams)

	if err != nil {
		return e.ErrorCaseHanding(e.NewBadInputError("Unknown field in body"))
	}

	err = gh.service.UpdateGuest(&bodyParams)

	if err != nil {
		return e.ErrorCaseHanding(err)
	}

	HandleJsonResponse(w, http.StatusOK, struct {
		Name string `json:"name"`
	}{Name: name})

	return nil
}

/**
 * Set a guest as left.
 * CURL CMD:  curl -X DELETE "localhost:3000/guests/<name>"
 */
func (gh *GuestHandler) DeleteGuest(w http.ResponseWriter, r *http.Request) *e.AppError {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	err := gh.service.DeleteGuest(name)

	if err != nil {
		return e.ErrorCaseHanding(err)
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}
