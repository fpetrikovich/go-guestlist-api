package exception

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

func ValidatePositiveInput(input int) error {
	if input < 0 {
		return &BadInputError{Input: fmt.Sprint(input)}
	}
	return nil
}

func ValidateStringInput(input string) error {
	if strings.Contains(input, " ") {
		return &BadInputError{Input: input}
	}
	return nil
}

func CheckDatabaseError(err error, id string, idType string, resource string) error {

	if err == nil {
		return nil
	} else if driverErr, ok := err.(*mysql.MySQLError); ok {
		switch driverErr.Number {
		case mysqlerr.ER_DUP_ENTRY:
			return NewAlreadyExistsError(id, idType, resource)
		}
	} else if err == sql.ErrNoRows {
		if len(id) > 0 {
			return NewNotFoundError(id, idType, resource) // user error
		} else {
			return NewMissingDataError(resource) // No id means there was no user input
		}
	}
	return err
}

func ErrorCaseHanding(err error) *AppError {
	switch err.(type) {
	case *NotFoundError:
		return &AppError{Error: err, Message: "[ERROR] " + err.Error(), Code: http.StatusNotFound}
	case *AlreadyExistsError:
		return &AppError{Error: err, Message: "[ERROR] " + err.Error(), Code: http.StatusConflict}
	case *BadInputError, *ExceedsCapacityError, *ArrivalStatusError:
		return &AppError{Error: err, Message: "[ERROR] " + err.Error(), Code: http.StatusBadRequest}
	case *MissingDataError:
		return &AppError{Error: err, Message: "[ERROR] " + err.Error(), Code: http.StatusInternalServerError}
	default:
		return &AppError{Error: err, Message: "[ERROR] Server error.", Code: http.StatusInternalServerError}
	}
}
