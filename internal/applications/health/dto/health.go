package dto

import "time"

type (
	HealthResponse struct {
		Message    string      `json:"message" validate:"required"`
		Status     string      `json:"status" validate:"required"`
		Timestamp  time.Time   `json:"timestamp" validate:"required"`
		Components *Components `json:"components"`
	}

	Components struct {
		Ctx   Details `json:"ctx" validate:"required"`
		Db    Details `json:"db" validate:"required"`
		Cache Details `json:"cache" validate:"required"`
		//other   Details    `json:"other" validate:"required"`
	}

	Details struct {
		Status string `json:"status" validate:"required"`
		Name   string `json:"name" validate:"required"`
	}
)
