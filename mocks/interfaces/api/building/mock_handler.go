// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/api/building/handler.go

// Package mock_building is a generated GoMock package.
package building

import (
	building "building_management/interfaces/api/building"
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockHandlerInterface is a mock of HandlerInterface interface.
type MockHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerInterfaceMockRecorder
}

// MockHandlerInterfaceMockRecorder is the mock recorder for MockHandlerInterface.
type MockHandlerInterfaceMockRecorder struct {
	mock *MockHandlerInterface
}

// NewMockHandlerInterface creates a new mock instance.
func NewMockHandlerInterface(ctrl *gomock.Controller) *MockHandlerInterface {
	mock := &MockHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandlerInterface) EXPECT() *MockHandlerInterfaceMockRecorder {
	return m.recorder
}

// GetCreateOrUpdateRequest mocks base method.
func (m *MockHandlerInterface) GetCreateOrUpdateRequest(c *fiber.Ctx) (building.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreateOrUpdateRequest", c)
	ret0, _ := ret[0].(building.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreateOrUpdateRequest indicates an expected call of GetCreateOrUpdateRequest.
func (mr *MockHandlerInterfaceMockRecorder) GetCreateOrUpdateRequest(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreateOrUpdateRequest", reflect.TypeOf((*MockHandlerInterface)(nil).GetCreateOrUpdateRequest), c)
}

// GetID mocks base method.
func (m *MockHandlerInterface) GetID(c *fiber.Ctx) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID", c)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetID indicates an expected call of GetID.
func (mr *MockHandlerInterfaceMockRecorder) GetID(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockHandlerInterface)(nil).GetID), c)
}
