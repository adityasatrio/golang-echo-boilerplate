package system_parameter

import "time"

//TODO delete
type SystemParameter struct {
	Key         string
	Value       string
	CreatedBy   string
	CreatedDate time.Time
	UpdatedBy   string
	UpdatedDate time.Time
}
