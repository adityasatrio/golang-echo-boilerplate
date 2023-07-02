package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/health/dto"
	"myapp/internal/applications/health/service"
	"myapp/internal/apputils"
	"time"
)

type HealthController struct {
	service service.HealthService
}

func NewHealthController(service service.HealthService) *HealthController {
	return &HealthController{
		service: service,
	}
}

func (controller *HealthController) Health(c echo.Context) error {

	queryFlag := c.QueryParam("flag")
	if queryFlag == "" {
		queryFlag = "default"
	}

	msgController := "hello from controller layer "
	result, err := controller.service.Health(c.Request().Context(), msgController, queryFlag)
	if err != nil {
		return err
	}

	var responseDto = dto.HealthResponse{
		Status:    "UP",
		Message:   result["final_msg"],
		Timestamp: time.Now(),
		Components: dto.Components{
			Ctx: dto.Details{
				Status: result["ctx_status"],
				Name:   result["ctx_name"],
			},
			Db: dto.Details{
				Status: result["db_status"],
				Name:   result["db_name"],
			},
		},
	}

	return apputils.Success(c, responseDto)
}
