package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fpetrikovich/go-guestlist/pkg/model"
	"github.com/fpetrikovich/go-guestlist/pkg/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GuestHandler_GetGuestList(t *testing.T) {
	name := "Flor"

	t.Run("Returns_OK_When_No_Errors", func(t *testing.T) {
		entourage := 10
		tableID := 1

		req, _ := http.NewRequest(http.MethodGet, "/guest_list", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIGuestService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetGuestList().
			Return([]model.GuestData{{Table: tableID, Name: name, Accompanying_guests: entourage}}, nil).
			Times(1)

		mh := NewGuestHandler(mockService)

		err := mh.GetGuestList(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var returnedGuests struct {
			Guest []model.GuestData `json:"guests"`
		}
		json.NewDecoder(rec.Body).Decode(&returnedGuests)

		assert.Nil(t, err)
		assert.Equal(t, tableID, returnedGuests.Guest[0].Table)
		assert.Equal(t, name, returnedGuests.Guest[0].Name)
		assert.Equal(t, entourage, returnedGuests.Guest[0].Accompanying_guests)
	})

	t.Run("Returns_ServerError_When_Service_Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/guest_list", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIGuestService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetGuestList().
			Return([]model.GuestData{}, errors.New("Error occurred")).
			Times(1)

		mh := NewGuestHandler(mockService)

		err := mh.GetGuestList(rec, req)

		assert.Equal(t, http.StatusInternalServerError, err.Code)
		assert.Equal(t, "[ERROR] Server error.", err.Message)
	})
}

func Test_GuestHandler_GetArrivedGuests(t *testing.T) {
	name := "Flor"

	t.Run("Returns_OK_When_No_Errors", func(t *testing.T) {
		entourage := 10
		time := time.Now().Format("2006-01-02 15:04:05")

		req, _ := http.NewRequest(http.MethodGet, "/guests", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIGuestService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetArrivedGuests().
			Return([]model.GuestArrival{{Name: name, Accompanying_guests: entourage, Arrived_at: time}}, nil).
			Times(1)

		mh := NewGuestHandler(mockService)

		err := mh.GetArrivedGuests(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var returnedGuests struct {
			Guest []model.GuestArrival `json:"guests"`
		}
		json.NewDecoder(rec.Body).Decode(&returnedGuests)

		assert.Nil(t, err)
		assert.Equal(t, name, returnedGuests.Guest[0].Name)
		assert.Equal(t, entourage, returnedGuests.Guest[0].Accompanying_guests)
		assert.Equal(t, time, returnedGuests.Guest[0].Arrived_at)
	})

	t.Run("Returns_ServerError_When_Service_Error", func(t *testing.T) {
		entourage := 10
		time := time.Now().Format("2006-01-02 15:04:05")

		req, _ := http.NewRequest(http.MethodGet, "/guests", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIGuestService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetArrivedGuests().
			Return([]model.GuestArrival{{Name: name, Accompanying_guests: entourage, Arrived_at: time}}, errors.New("Unknown error.")).
			Times(1)

		mh := NewGuestHandler(mockService)

		err := mh.GetArrivedGuests(rec, req)

		assert.Equal(t, http.StatusInternalServerError, err.Code)
		assert.Equal(t, "[ERROR] Server error.", err.Message)
	})
}
