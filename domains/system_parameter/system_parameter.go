package system_parameter

import "time"

type SystemParameter struct {
	Key         string
	Value       string
	CreatedBy   string
	CreatedDate time.Time
	UpdatedBy   string
	UpdatedDate time.Time
}
