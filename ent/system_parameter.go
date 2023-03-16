// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"myapp/ent/system_parameter"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// System_parameter is the model entity for the System_parameter schema.
type System_parameter struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*System_parameter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case system_parameter.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case system_parameter.FieldID:
			values[i] = new(sql.NullInt64)
		case system_parameter.FieldKey, system_parameter.FieldValue, system_parameter.FieldCreatedBy, system_parameter.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case system_parameter.FieldCreatedAt, system_parameter.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type System_parameter", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the System_parameter fields.
func (sp *System_parameter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case system_parameter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = int(value.Int64)
		case system_parameter.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				sp.Key = value.String
			}
		case system_parameter.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				sp.Value = value.String
			}
		case system_parameter.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				sp.IsDeleted = value.Bool
			}
		case system_parameter.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				sp.CreatedBy = value.String
			}
		case system_parameter.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = value.Time
			}
		case system_parameter.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sp.UpdatedBy = value.String
			}
		case system_parameter.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this System_parameter.
// Note that you need to call System_parameter.Unwrap() before calling this method if this System_parameter
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *System_parameter) Update() *SystemParameterUpdateOne {
	return (&System_parameterClient{config: sp.config}).UpdateOne(sp)
}

// Unwrap unwraps the System_parameter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *System_parameter) Unwrap() *System_parameter {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: System_parameter is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *System_parameter) String() string {
	var builder strings.Builder
	builder.WriteString("System_parameter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("key=")
	builder.WriteString(sp.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(sp.Value)
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", sp.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(sp.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(sp.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sp.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// System_parameters is a parsable slice of System_parameter.
type System_parameters []*System_parameter

func (sp System_parameters) config(cfg config) {
	for _i := range sp {
		sp[_i].config = cfg
	}
}