// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/api/apartment/service.go

// Package mock_apartment is a generated GoMock package.
package apartment

import (
	apartment "building_management/interfaces/api/apartment"
	models "building_management/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceInterface is a mock of ServiceInterface interface.
type MockServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInterfaceMockRecorder
}

// MockServiceInterfaceMockRecorder is the mock recorder for MockServiceInterface.
type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

// NewMockServiceInterface creates a new mock instance.
func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateOrUpdateApartment mocks base method.
func (m *MockServiceInterface) CreateOrUpdateApartment(ctx context.Context, apartment apartment.ApartmentRequest) (*models.Apartment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateApartment", ctx, apartment)
	ret0, _ := ret[0].(*models.Apartment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateApartment indicates an expected call of CreateOrUpdateApartment.
func (mr *MockServiceInterfaceMockRecorder) CreateOrUpdateApartment(ctx, apartment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateApartment", reflect.TypeOf((*MockServiceInterface)(nil).CreateOrUpdateApartment), ctx, apartment)
}

// DeleteApartment mocks base method.
func (m *MockServiceInterface) DeleteApartment(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApartment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApartment indicates an expected call of DeleteApartment.
func (mr *MockServiceInterfaceMockRecorder) DeleteApartment(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApartment", reflect.TypeOf((*MockServiceInterface)(nil).DeleteApartment), ctx, id)
}

// GetApartmentByID mocks base method.
func (m *MockServiceInterface) GetApartmentByID(ctx context.Context, id int) (*models.Apartment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApartmentByID", ctx, id)
	ret0, _ := ret[0].(*models.Apartment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApartmentByID indicates an expected call of GetApartmentByID.
func (mr *MockServiceInterfaceMockRecorder) GetApartmentByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApartmentByID", reflect.TypeOf((*MockServiceInterface)(nil).GetApartmentByID), ctx, id)
}

// GetApartments mocks base method.
func (m *MockServiceInterface) GetApartments(ctx context.Context) (models.ApartmentSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApartments", ctx)
	ret0, _ := ret[0].(models.ApartmentSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApartments indicates an expected call of GetApartments.
func (mr *MockServiceInterfaceMockRecorder) GetApartments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApartments", reflect.TypeOf((*MockServiceInterface)(nil).GetApartments), ctx)
}

// GetApartmentsByBuilding mocks base method.
func (m *MockServiceInterface) GetApartmentsByBuilding(ctx context.Context, buildingID int) (models.ApartmentSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApartmentsByBuilding", ctx, buildingID)
	ret0, _ := ret[0].(models.ApartmentSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApartmentsByBuilding indicates an expected call of GetApartmentsByBuilding.
func (mr *MockServiceInterfaceMockRecorder) GetApartmentsByBuilding(ctx, buildingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApartmentsByBuilding", reflect.TypeOf((*MockServiceInterface)(nil).GetApartmentsByBuilding), ctx, buildingID)
}
