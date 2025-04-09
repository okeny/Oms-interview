package apartment

import (
	"building_management/interfaces/api/apartment"
	"building_management/models"
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// Import the mocks with correct package path
	mock "building_management/mocks/interfaces/api/apartment"
)

func SetUp(ctrl *gomock.Controller) (Controller,
	*fiber.App,
	*mock.MockServiceInterface,
	*mock.MockHandlerInterface) {

	mockHandler := mock.NewMockHandlerInterface(ctrl)
	mockService := mock.NewMockServiceInterface(ctrl)
	// Pass mocks as Handler and Service interfaces
	controller := NewController(mockHandler, mockService)

	app := fiber.New()
	return controller, app, mockService, mockHandler
}

func TestController_GetApartments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, _ := SetUp(ctrl)
	app.Get("/apartments", controller.GetApartments)

	tests := []struct {
		name         string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - apartments retrieved",
			mockSetup: func() {
				mockService.EXPECT().GetApartments(gomock.Any()).Return(models.ApartmentSlice{
					&models.Apartment{ID: 1, BuildingID: 1, Number: "A1"},
				}, nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `[{"id":1,"building_id":1,"number":"A1","floor":0,"sq_meters":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name: "Error - service fails",
			mockSetup: func() {
				mockService.EXPECT().GetApartments(gomock.Any()).Return(nil, errors.New("db error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"db error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("GET", "/apartments", nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)
			body, _ := json.MarshalIndent(json.RawMessage(tt.expectedBody), "", "  ")
			assert.JSONEq(t, string(body), string(bodyBytes))
		})
	}
}

func TestController_GetApartmentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, mockHandler := SetUp(ctrl)
	app.Get("/apartments/:id", controller.GetApartmentByID)

	tests := []struct {
		name         string
		url          string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - apartment found",
			url:  "/apartments/1",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(1, nil)
				mockService.EXPECT().GetApartmentByID(gomock.Any(), 1).Return(&models.Apartment{
					ID:         1,
					BuildingID: 1,
					Number:     "A1",
					Floor:      0,
					SQMeters:   0,
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
				}, nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `{
				"id": 1,
				"building_id": 1,
				"number": "A1",
				"floor": 0,
				"sq_meters": 0,
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			}`,
		},
		{
			name: "Error - invalid ID",
			url:  "/apartments/invalid",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(0, errors.New("invalid id"))
			},
			expectedCode: fiber.StatusBadRequest,
			expectedBody: `{"error":"Invalid apartment ID"}`,
		},
		{
			name: "Error - service fails",
			url:  "/apartments/2",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(2, nil)
				mockService.EXPECT().GetApartmentByID(gomock.Any(), 2).Return(nil, errors.New("not found"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("GET", tt.url, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			bodyBytes, _ := io.ReadAll(resp.Body)
			body, _ := json.MarshalIndent(json.RawMessage(tt.expectedBody), "", "  ")
			assert.JSONEq(t, string(body), string(bodyBytes))
		})
	}
}
func TestController_GetApartmentsByBuilding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, mockHandler := SetUp(ctrl)
	app.Get("/buildings/:id/apartments", controller.GetApartmentsByBuilding)

	tests := []struct {
		name         string
		url          string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - apartments found",
			url:  "/buildings/1/apartments",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(1, nil) // Ensure GetID is mocked correctly
				mockService.EXPECT().GetApartmentsByBuilding(gomock.Any(), 1).Return([]*models.Apartment{
					{ID: 1, BuildingID: 1, Number: "A1", Floor: 0, SQMeters: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
				}, nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `[
				{
					"id": 1,
					"building_id": 1,
					"number": "A1",
					"floor": 0,
					"sq_meters": 0,
					"created_at": "0001-01-01T00:00:00Z",
					"updated_at": "0001-01-01T00:00:00Z"
				}
			]`,
		},
		{
			name: "Error - no apartments found",
			url:  "/buildings/2/apartments",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(2, nil) // Ensure GetID is mocked for this test case as well
				mockService.EXPECT().GetApartmentsByBuilding(gomock.Any(), 2).Return(nil, nil) // No apartments found
			},
			expectedCode: fiber.StatusNotFound,
			expectedBody: `{"error":"apartments not found"}`,
		},
		{
			name: "Error - service fails",
			url:  "/buildings/1/apartments",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(1, nil) // Mock GetID call
				mockService.EXPECT().GetApartmentsByBuilding(gomock.Any(), 1).Return(nil, errors.New("internal server error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("GET", tt.url, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			bodyBytes, _ := io.ReadAll(resp.Body)
			body, _ := json.MarshalIndent(json.RawMessage(tt.expectedBody), "", "  ")
			assert.JSONEq(t, string(body), string(bodyBytes))
		})
	}
}

func TestController_CreateOrUpdateApartment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, mockHandler := SetUp(ctrl)
	app.Post("/apartments", controller.CreateOrUpdateApartment)

	tests := []struct {
		name         string
		body         string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - apartment created",
			body: `{"id":0,"building_id":1,"number":"A1","floor":1,"sq_meters":50}`,
			mockSetup: func() {
				req := apartment.ApartmentRequest{ID: 0, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50}
				mockHandler.EXPECT().GetCreateOrUpdateRequest(gomock.Any()).Return(req, nil)
				mockService.EXPECT().CreateOrUpdateApartment(gomock.Any(), req).
					Return(&models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50}, nil)
			},
			expectedCode: fiber.StatusCreated,
			expectedBody: `{
				"id": 1,
				"building_id": 1,
				"number": "A1",
				"floor": 1,
				"sq_meters": 50,
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			}`,
		},
		{
			name: "Error - invalid request",
			body: `{"id":"invalid"}`,
			mockSetup: func() {
				mockHandler.EXPECT().GetCreateOrUpdateRequest(gomock.Any()).Return(apartment.ApartmentRequest{}, errors.New("invalid request"))
			},
			expectedCode: fiber.StatusBadRequest,
			expectedBody: `{"error":"invalid request"}`,
		},
		{
			name: "Error - service fails",
			body: `{"id":0,"building_id":1,"number":"A1"}`,
			mockSetup: func() {
				req := apartment.ApartmentRequest{ID: 0, BuildingID: 1, Number: "A1"}
				mockHandler.EXPECT().GetCreateOrUpdateRequest(gomock.Any()).Return(req, nil)
				mockService.EXPECT().CreateOrUpdateApartment(gomock.Any(), req).
					Return(&models.Apartment{}, errors.New("upsert error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"upsert error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("POST", "/apartments", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)
			body, _ := json.MarshalIndent(json.RawMessage(tt.expectedBody), "", "  ")
			assert.JSONEq(t, string(body), string(bodyBytes))
		})
	}
}

func TestController_DeleteApartment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, mockHandler := SetUp(ctrl)
	app.Delete("/apartments/:id", controller.DeleteApartment)

	tests := []struct {
		name         string
		url          string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - apartment deleted",
			url:  "/apartments/1",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(1, nil)
				mockService.EXPECT().DeleteApartment(gomock.Any(), 1).Return(nil)
			},
			expectedCode: fiber.StatusNoContent,
			expectedBody: ``,
		},
		{
			name: "Error - invalid ID",
			url:  "/apartments/invalid",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(0, errors.New("invalid id"))
			},
			expectedCode: fiber.StatusBadRequest,
			expectedBody: `{"error":"Invalid apartment ID"}`,
		},
		{
			name: "Error - service fails",
			url:  "/apartments/2",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(2, nil)
				mockService.EXPECT().DeleteApartment(gomock.Any(), 2).Return(errors.New("delete error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"delete error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("DELETE", tt.url, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)
			if tt.expectedBody != "" {
				body, _ := json.MarshalIndent(json.RawMessage(tt.expectedBody), "", "  ")
				assert.JSONEq(t, string(body), string(bodyBytes))
			} else {
				assert.Empty(t, string(bodyBytes))
			}
		})
	}
}
