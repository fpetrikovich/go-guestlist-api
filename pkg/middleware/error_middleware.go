package middleware

import (
	"log"
	"net/http"

	e "github.com/fpetrikovich/go-guestlist/pkg/exception"
)

/*
Type `AppHandler` which is a function type that takes an `http.ResponseWriter` and
an `*http.Request` as input and returns a pointer to an e.AppError struct.
*/
type AppHandler func(http.ResponseWriter, *http.Request) *e.AppError

/*
`ServeHTTP` takes an `http.ResponseWriter` and an `*http.Request` as input.
If a non-nil error is returned by calling `fn` with the `http.ResponseWriter` and `*http.Request` as arguments,
the error's details are logged using the log package and an HTTP error response is sent
using the `http.Error` function with the error message and code specified in the `AppError` struct.
*/
func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *AppError, not os.Error.
		log.Print("[ERROR] ", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}
