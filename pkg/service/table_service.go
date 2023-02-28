package service

import (
	"github.com/fpetrikovich/go-guestlist/pkg/model"
	"github.com/fpetrikovich/go-guestlist/pkg/repository"
)

/*
The purpose of this code is to define a Go service for managing event tables.

It implements a `DefaultEventTableService` struct that has a field for an `IEventTableRepository` interface.
The code provides functions for retrieving information about event tables
(e.g. `GetTables()`, `GetTable(id int)`, `GetEmptySeats()`, `GetEmptySeatsAtTable(id int)`) and
also for creating and deleting event tables (e.g. `CreateTable(*model.EventTable)`, `DeleteTable(id int)`).

The functions interact with the IEventTableRepository to perform the desired operations.
*/
type DefaultEventTableService struct {
	tableRepository repository.IEventTableRepository
}

func NewDefaultEventTableService(tRepo repository.IEventTableRepository) *DefaultEventTableService {
	return &DefaultEventTableService{
		tableRepository: tRepo,
	}
}

func (d *DefaultEventTableService) GetTables() ([]model.EventTable, error) {
	return d.tableRepository.GetTables()
}

func (d *DefaultEventTableService) GetTable(id int) (*model.EventTable, error) {
	return d.tableRepository.GetTable(id)
}

func (d *DefaultEventTableService) CreateTable(table *model.EventTable) (*model.EventTable, error) {
	return d.tableRepository.CreateTable(table)
}

func (d *DefaultEventTableService) DeleteTable(id int) error {
	return nil
}

func (d *DefaultEventTableService) GetEmptySeatsAtTable(id int) (int, error) {
	return d.tableRepository.GetEmptySeatsAtTable(id)
}

func (d *DefaultEventTableService) GetEmptySeats() (int, error) {
	return d.tableRepository.GetEmptySeats()
}
