package apartment

import (
	"building_management/interfaces/api/apartment"
	"building_management/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) Handler {
	return Handler{db: db}
}

func (h Handler) GetID(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (h Handler) GetCreateOrUpdateRequest(c *fiber.Ctx) (apartment.Request, error) {
	var request apartment.Request
	// Parse request body into apartment struct
	if err := c.BodyParser(&request); err != nil {
		return apartment.Request{}, err
	}
	_, err := h.CheckBuilding(c.Context(), request.BuildingID)
	if err != nil {
		return apartment.Request{}, err
	}

	return request, nil
}

func (h Handler) CheckBuilding(ctx context.Context, buildingID int) (bool, error) {
	_, err := models.Buildings(models.BuildingWhere.ID.EQ(buildingID)).One(ctx, h.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("building with id %d does not exist", buildingID)
		}
		return false, fmt.Errorf("failed to check building: %w", err)
	}

	return true, nil
}
