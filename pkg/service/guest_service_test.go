package service

import (
	"testing"

	ex "github.com/fpetrikovich/go-guestlist/pkg/exception"
	"github.com/fpetrikovich/go-guestlist/pkg/model"
	"github.com/fpetrikovich/go-guestlist/pkg/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_DefaultTableService_UpdateGuest(t *testing.T) {
	name := "Flor"
	guest := model.Guest{
		GuestID:       1,
		Name:          name,
		Entourage:     3,
		ArrivalStatus: model.GuestStatus("not_arrived"),
	}

	t.Run("Return_BadInput_When_Entourage_Is_Negative", func(t *testing.T) {

		testCase := model.GuestData{
			Name:                name,
			Accompanying_guests: -4,
			Table:               1,
		}

		dms := NewDefaultGuestService(nil, nil)
		err := dms.UpdateGuest(&testCase)
		assert.Equal(t, err.Error(), ex.NewBadInputError("-4").Error())
	})

	t.Run("Return_NotFound_When_Guest_Doesnt_Exist", func(t *testing.T) {

		testCase := model.GuestData{
			Name:                name,
			Accompanying_guests: 4,
			Table:               1,
		}

		errNotFound := ex.NewNotFoundError(name, "name", "guest")

		mockRepository := repository.NewMockIGuestRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().
			GetGuest(name).
			Return(&model.Guest{}, errNotFound).
			Times(1)

		ms := NewDefaultGuestService(mockRepository, nil)

		err := ms.UpdateGuest(&testCase)
		assert.Equal(t, err.Error(), errNotFound.Error())
	})

	t.Run("Set_Rejected_When_Entourage_Exceeds_Capacity", func(t *testing.T) {

		testCase := model.GuestData{
			Name:                name,
			Accompanying_guests: 7,
			Table:               1,
		}

		mockRepository := repository.NewMockIGuestRepository(gomock.NewController(t))

		mockRepository.
			EXPECT().
			GetGuest(name).
			Return(&guest, nil).
			Times(1)

		mockRepository.
			EXPECT().
			GetGuestTableFreeSeats(name).
			Return(3, nil).
			Times(1)

		mockRepository.
			EXPECT().
			UpdateGuest(&guest).
			Return(nil).
			Times(1)
		ms := NewDefaultGuestService(mockRepository, nil)

		_ = ms.UpdateGuest(&testCase)
		assert.Equal(t, guest.ArrivalStatus, model.GuestStatus("rejected"))
	})

	t.Run("Set_Arrived_When_Entourage_Doesnt_Exceeds_Capacity", func(t *testing.T) {

		testCases := []model.GuestData{{
			// same amount
			Name:                name,
			Accompanying_guests: 3,
			Table:               1,
		}, {
			// More guests
			Name:                name,
			Accompanying_guests: 6,
			Table:               1,
		}, {
			// Less guests
			Name:                name,
			Accompanying_guests: 2,
			Table:               1,
		}}

		mockRepository := repository.NewMockIGuestRepository(gomock.NewController(t))

		mockRepository.
			EXPECT().
			GetGuest(name).
			Return(&guest, nil).
			Times(len(testCases))

		mockRepository.
			EXPECT().
			GetGuestTableFreeSeats(name).
			Return(3, nil).
			Times(len(testCases))

		mockRepository.
			EXPECT().
			UpdateGuest(&guest).
			Return(nil).
			Times(len(testCases))

		ms := NewDefaultGuestService(mockRepository, nil)

		for _, test := range testCases {
			err := ms.UpdateGuest(&test)
			assert.Equal(t, guest.ArrivalStatus, model.GuestStatus("arrived"))
			assert.Nil(t, err)
		}
	})
}

func Test_DefaultGuestService_CreateGuest(t *testing.T) {
	name := "Flor"
	_ = model.Guest{
		GuestID:       1,
		Name:          name,
		Entourage:     3,
		ArrivalStatus: model.GuestStatus("not_arrived"),
	}
	t.Run("Return_BadRequest_When_Invalid_Entourage", func(t *testing.T) {
		testCase := model.GuestData{
			Name:                name,
			Accompanying_guests: -4,
			Table:               1,
		}

		dms := NewDefaultGuestService(nil, nil)
		err := dms.CreateGuest(&testCase)
		assert.Equal(t, err.Error(), ex.NewBadInputError("-4").Error())
	})
	t.Run("Return_CapacityError_When_Entourage_Exceed_Capacity", func(t *testing.T) {
		testCase := model.GuestData{
			Name:                name,
			Accompanying_guests: 4,
			Table:               1,
		}

		mockTableService := NewMockIEventTableService(gomock.NewController(t))

		mockTableService.
			EXPECT().
			GetEmptySeatsAtTable(testCase.Table).
			Return(4, nil).
			Times(1)

		ms := NewDefaultGuestService(nil, mockTableService)
		err := ms.CreateGuest(&testCase)

		assert.Equal(t, err.Error(), ex.NewExceedsCapacityError(4, 1).Error())
	})

	t.Run("Return_AlreadyExists_When_Duplicate_Name", func(t *testing.T) {
		testCase := model.GuestData{
			Name:                name,
			Accompanying_guests: 4,
			Table:               1,
		}

		mockTableService := NewMockIEventTableService(gomock.NewController(t))

		mockTableService.
			EXPECT().
			GetEmptySeatsAtTable(testCase.Table).
			Return(10, nil).
			Times(1)

		mockRepository := repository.NewMockIGuestRepository(gomock.NewController(t))

		mockRepository.
			EXPECT().
			CreateGuest(&testCase).
			Return(ex.NewAlreadyExistsError(name, "name", "guest")).
			Times(1)

		ms := NewDefaultGuestService(mockRepository, mockTableService)
		err := ms.CreateGuest(&testCase)
		assert.Equal(t, err.Error(), ex.NewAlreadyExistsError(name, "name", "guest").Error())
	})

	t.Run("Return_Success_When_Valid_Guest", func(t *testing.T) {
		testCase := model.GuestData{
			Name:                name,
			Accompanying_guests: 4,
			Table:               1,
		}

		mockTableService := NewMockIEventTableService(gomock.NewController(t))

		mockTableService.
			EXPECT().
			GetEmptySeatsAtTable(testCase.Table).
			Return(10, nil).
			Times(1)

		mockRepository := repository.NewMockIGuestRepository(gomock.NewController(t))

		mockRepository.
			EXPECT().
			CreateGuest(&testCase).
			Return(nil).
			Times(1)

		ms := NewDefaultGuestService(mockRepository, mockTableService)
		err := ms.CreateGuest(&testCase)
		assert.Nil(t, err)
	})

}
