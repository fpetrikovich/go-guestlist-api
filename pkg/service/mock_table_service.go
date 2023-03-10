// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/table_service_interface.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	model "github.com/fpetrikovich/go-guestlist/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIEventTableService is a mock of IEventTableService interface.
type MockIEventTableService struct {
	ctrl     *gomock.Controller
	recorder *MockIEventTableServiceMockRecorder
}

// MockIEventTableServiceMockRecorder is the mock recorder for MockIEventTableService.
type MockIEventTableServiceMockRecorder struct {
	mock *MockIEventTableService
}

// NewMockIEventTableService creates a new mock instance.
func NewMockIEventTableService(ctrl *gomock.Controller) *MockIEventTableService {
	mock := &MockIEventTableService{ctrl: ctrl}
	mock.recorder = &MockIEventTableServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventTableService) EXPECT() *MockIEventTableServiceMockRecorder {
	return m.recorder
}

// CreateTable mocks base method.
func (m *MockIEventTableService) CreateTable(table *model.EventTable) (*model.EventTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTable", table)
	ret0, _ := ret[0].(*model.EventTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockIEventTableServiceMockRecorder) CreateTable(table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockIEventTableService)(nil).CreateTable), table)
}

// DeleteTable mocks base method.
func (m *MockIEventTableService) DeleteTable(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTable", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTable indicates an expected call of DeleteTable.
func (mr *MockIEventTableServiceMockRecorder) DeleteTable(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTable", reflect.TypeOf((*MockIEventTableService)(nil).DeleteTable), id)
}

// GetEmptySeats mocks base method.
func (m *MockIEventTableService) GetEmptySeats() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmptySeats")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmptySeats indicates an expected call of GetEmptySeats.
func (mr *MockIEventTableServiceMockRecorder) GetEmptySeats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmptySeats", reflect.TypeOf((*MockIEventTableService)(nil).GetEmptySeats))
}

// GetEmptySeatsAtTable mocks base method.
func (m *MockIEventTableService) GetEmptySeatsAtTable(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmptySeatsAtTable", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmptySeatsAtTable indicates an expected call of GetEmptySeatsAtTable.
func (mr *MockIEventTableServiceMockRecorder) GetEmptySeatsAtTable(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmptySeatsAtTable", reflect.TypeOf((*MockIEventTableService)(nil).GetEmptySeatsAtTable), id)
}

// GetTable mocks base method.
func (m *MockIEventTableService) GetTable(id int) (*model.EventTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTable", id)
	ret0, _ := ret[0].(*model.EventTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTable indicates an expected call of GetTable.
func (mr *MockIEventTableServiceMockRecorder) GetTable(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTable", reflect.TypeOf((*MockIEventTableService)(nil).GetTable), id)
}

// GetTables mocks base method.
func (m *MockIEventTableService) GetTables() ([]model.EventTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTables")
	ret0, _ := ret[0].([]model.EventTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTables indicates an expected call of GetTables.
func (mr *MockIEventTableServiceMockRecorder) GetTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTables", reflect.TypeOf((*MockIEventTableService)(nil).GetTables))
}
