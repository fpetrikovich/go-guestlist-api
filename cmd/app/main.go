package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fpetrikovich/go-guestlist/pkg/handler"
	mw "github.com/fpetrikovich/go-guestlist/pkg/middleware"
	"github.com/fpetrikovich/go-guestlist/pkg/repository"
	"github.com/fpetrikovich/go-guestlist/pkg/service"
)

func main() {
	router := mux.NewRouter()

	log.Print("[INFO] Server is up!")

	// Connect to the database
	dbRepository := repository.NewMySQLRepository()
	defer dbRepository.Connection.Close()

	initRoutes(router, dbRepository)

	err := http.ListenAndServe(":3000", router)

	log.Fatal(err)
}

/*
The initRoutes function sets up HTTP routes for a `mux.Router` using a `repository.MySQLRepository` for database access.
It takes in a `mux.Router` pointer and a `repository.MySQLRepository` pointer as parameters and maps URL paths to their respective handlers.
This function provides a centralized location for managing application routes.
*/
func initRoutes(router *mux.Router, dbRepo *repository.MySQLRepository) {

	// Create handlers
	tableHandler, guestHandler := createHandlers(dbRepo.Connection)

	// Table Routes
	router.Handle("/tables/{id}", mw.AppHandler(tableHandler.GetTable)).Methods("GET")
	router.Handle("/tables", mw.AppHandler(tableHandler.GetTables)).Methods("GET")
	router.Handle("/tables", mw.AppHandler(tableHandler.CreateTable)).Methods("POST")
	router.Handle("/seats_empty", mw.AppHandler(tableHandler.GetEmptySeats)).Methods("GET")
	// Guest Routes
	router.Handle("/guest_list/{name}", mw.AppHandler(guestHandler.GetGuest)).Methods("GET")
	router.Handle("/guest_list", mw.AppHandler(guestHandler.GetGuestList)).Methods("GET")
	router.Handle("/guest_list/{name}", mw.AppHandler(guestHandler.CreateGuest)).Methods("POST")
	router.Handle("/guests/{name}", mw.AppHandler(guestHandler.UpdateGuest)).Methods("PUT")
	router.Handle("/guests", mw.AppHandler(guestHandler.GetArrivedGuests)).Methods("GET")
	router.Handle("/guests/{name}", mw.AppHandler(guestHandler.DeleteGuest)).Methods("DELETE")

	// ping
	router.HandleFunc("/ping", handlerPing)
}

/*
The `createHandlers` function creates two handlers, `handler.EventTableHandler` and `handler.GuestHandler`,
for a SQL database represented by the `sql.DB` pointer `con`. It returns two pointers to these handlers.
The purpose of this function is to create instances of the event table and guest handlers
and pass in the database connection so they can access the database.
*/
func createHandlers(con *sql.DB) (*handler.EventTableHandler, *handler.GuestHandler) {
	// Table
	tableRepository := repository.NewMySQLEventTableRepository(con)
	tableService := service.NewDefaultEventTableService(tableRepository)
	// Guest
	guestRepository := repository.NewMySQLGuestRepository(con)
	guestService := service.NewDefaultGuestService(guestRepository, tableService)
	// Handlers
	return handler.NewEventTableHandler(tableService), handler.NewGuestHandler(guestService)
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}
