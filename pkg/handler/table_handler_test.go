package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fpetrikovich/go-guestlist/pkg/model"
	"github.com/fpetrikovich/go-guestlist/pkg/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_TableHandler_GetTables(t *testing.T) {
	t.Run("Returns_OK_When_No_Errors", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tables", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIEventTableService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetTables().
			Return([]model.EventTable{{TableID: 1, Capacity: 10}}, nil).
			Times(1)

		mh := NewEventTableHandler(mockService)

		mh.GetTables(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var returnedTables struct {
			Tables []model.EventTable `json:"tables"`
		}
		json.NewDecoder(rec.Body).Decode(&returnedTables)

		assert.Equal(t, 1, returnedTables.Tables[0].TableID)
		assert.Equal(t, 10, returnedTables.Tables[0].Capacity)
	})

	t.Run("Returns_ServerError_When_Service_Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tables", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIEventTableService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetTables().
			Return([]model.EventTable{{TableID: 1, Capacity: 10}}, errors.New("Error occurred")).
			Times(1)

		mh := NewEventTableHandler(mockService)

		err := mh.GetTables(rec, req)

		assert.Equal(t, http.StatusInternalServerError, err.Code)
	})
}
