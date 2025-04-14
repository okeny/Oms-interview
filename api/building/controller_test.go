package building

import (
	"building_management/interfaces/api/building"
	"building_management/models"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// Import the mocks with the correct package path
	mock "building_management/mocks/interfaces/api/building"
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

func TestController_GetBuildings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, _ := SetUp(ctrl)
	app.Get("/buildings", controller.GetBuildings)

	tests := []struct {
		name         string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - buildings retrieved",
			mockSetup: func() {
				mockService.EXPECT().GetBuildings(gomock.Any()).Return([]*models.Building{
					{ID: 1, Name: "Building A"},
				}, nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `[{
				"id": 1,
				"name": "Building A",
				"address": "",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			}]`,
		},
		{
			name: "Error - service fails",
			mockSetup: func() {
				mockService.EXPECT().GetBuildings(gomock.Any()).Return(nil, errors.New("db error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"db error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("GET", "/buildings", http.NoBody)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(bodyBytes))
		})
	}
}

func TestController_GetBuildingByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, mockHandler := SetUp(ctrl)
	app.Get("/buildings/:id", controller.GetBuildingByID)

	tests := []struct {
		name         string
		url          string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - building found",
			url:  "/buildings/1",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(1, nil)
				mockService.EXPECT().GetBuildingByID(gomock.Any(), 1).Return(&models.Building{
					ID:   1,
					Name: "Building A",
				}, nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `{
				"id": 1,
				"name": "Building A",
				"address": "",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			}`,
		},
		{
			name: "Error - invalid ID",
			url:  "/buildings/invalid",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(0, errors.New("invalid id"))
			},
			expectedCode: fiber.StatusBadRequest,
			expectedBody: `{"error":"Invalid building ID"}`,
		},
		{
			name: "Error - building not found",
			url:  "/buildings/2",
			mockSetup: func() {
				mockHandler.EXPECT().GetID(gomock.Any()).Return(2, nil)
				mockService.EXPECT().GetBuildingByID(gomock.Any(), 2).Return(nil, errors.New("not found"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("GET", tt.url, http.NoBody)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(bodyBytes))
		})
	}
}

func TestController_CreateOrUpdateBuilding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, mockHandler := SetUp(ctrl)
	app.Post("/buildings", controller.CreateOrUpdateBuilding)

	tests := []struct {
		name         string
		body         string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - building created",
			body: `{"name":"Building A"}`,
			mockSetup: func() {
				req := building.Request{Name: "Building A"}
				mockHandler.EXPECT().GetCreateOrUpdateRequest(gomock.Any()).Return(req, nil)
				mockService.EXPECT().CreateOrUpdateBuilding(gomock.Any(), req).
					Return(&models.Building{ID: 1, Name: "Building A"}, nil)
			},
			expectedCode: fiber.StatusCreated,
			expectedBody: `{
				"id": 1,
				"name": "Building A",
				"address":"",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			}`,
		},
		{
			name: "Error - invalid request",
			body: `{"name":""}`,
			mockSetup: func() {
				mockHandler.EXPECT().GetCreateOrUpdateRequest(gomock.Any()).Return(building.Request{}, errors.New("invalid request"))
			},
			expectedCode: fiber.StatusBadRequest,
			expectedBody: `{"error":"invalid request"}`,
		},
		{
			name: "Error - service fails",
			body: `{"name":"Building A"}`,
			mockSetup: func() {
				req := building.Request{Name: "Building A"}
				mockHandler.EXPECT().GetCreateOrUpdateRequest(gomock.Any()).Return(req, nil)
				mockService.EXPECT().CreateOrUpdateBuilding(gomock.Any(), req).
					Return(nil, errors.New("upsert error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"error":"upsert error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("POST", "/buildings", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(bodyBytes))
		})
	}
}

func TestController_DeleteBuilding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller, app, mockService, _ := SetUp(ctrl)
	app.Delete("/buildings/:id", controller.DeleteBuilding)

	tests := []struct {
		name         string
		mockSetup    func()
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success - building deleted",
			mockSetup: func() {
				mockService.EXPECT().DeleteBuilding(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `{"message":"Building deleted"}`, // Make sure this matches the response
		},
		{
			name: "Error - building not found",
			mockSetup: func() {
				mockService.EXPECT().DeleteBuilding(gomock.Any(), gomock.Any()).Return(errors.New("building not found"))
			},
			expectedCode: fiber.StatusNotFound,
			expectedBody: `{"error":"building not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("DELETE", "/buildings/1", http.NoBody)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			// Check response status code
			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			// Read the response body
			bodyBytes, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(bodyBytes))
		})
	}
}
