// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myapp/ent/predicate"
	"myapp/ent/userdevice"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserDeviceUpdate is the builder for updating UserDevice entities.
type UserDeviceUpdate struct {
	config
	hooks    []Hook
	mutation *UserDeviceMutation
}

// Where appends a list predicates to the UserDeviceUpdate builder.
func (udu *UserDeviceUpdate) Where(ps ...predicate.UserDevice) *UserDeviceUpdate {
	udu.mutation.Where(ps...)
	return udu
}

// SetUserID sets the "user_id" field.
func (udu *UserDeviceUpdate) SetUserID(u uint64) *UserDeviceUpdate {
	udu.mutation.ResetUserID()
	udu.mutation.SetUserID(u)
	return udu
}

// AddUserID adds u to the "user_id" field.
func (udu *UserDeviceUpdate) AddUserID(u int64) *UserDeviceUpdate {
	udu.mutation.AddUserID(u)
	return udu
}

// SetVersion sets the "version" field.
func (udu *UserDeviceUpdate) SetVersion(s string) *UserDeviceUpdate {
	udu.mutation.SetVersion(s)
	return udu
}

// SetPlatform sets the "platform" field.
func (udu *UserDeviceUpdate) SetPlatform(s string) *UserDeviceUpdate {
	udu.mutation.SetPlatform(s)
	return udu
}

// SetLatestSkipUpdate sets the "latest_skip_update" field.
func (udu *UserDeviceUpdate) SetLatestSkipUpdate(t time.Time) *UserDeviceUpdate {
	udu.mutation.SetLatestSkipUpdate(t)
	return udu
}

// SetNillableLatestSkipUpdate sets the "latest_skip_update" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableLatestSkipUpdate(t *time.Time) *UserDeviceUpdate {
	if t != nil {
		udu.SetLatestSkipUpdate(*t)
	}
	return udu
}

// ClearLatestSkipUpdate clears the value of the "latest_skip_update" field.
func (udu *UserDeviceUpdate) ClearLatestSkipUpdate() *UserDeviceUpdate {
	udu.mutation.ClearLatestSkipUpdate()
	return udu
}

// SetCreatedAt sets the "created_at" field.
func (udu *UserDeviceUpdate) SetCreatedAt(t time.Time) *UserDeviceUpdate {
	udu.mutation.SetCreatedAt(t)
	return udu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableCreatedAt(t *time.Time) *UserDeviceUpdate {
	if t != nil {
		udu.SetCreatedAt(*t)
	}
	return udu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (udu *UserDeviceUpdate) ClearCreatedAt() *UserDeviceUpdate {
	udu.mutation.ClearCreatedAt()
	return udu
}

// SetUpdatedAt sets the "updated_at" field.
func (udu *UserDeviceUpdate) SetUpdatedAt(t time.Time) *UserDeviceUpdate {
	udu.mutation.SetUpdatedAt(t)
	return udu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableUpdatedAt(t *time.Time) *UserDeviceUpdate {
	if t != nil {
		udu.SetUpdatedAt(*t)
	}
	return udu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (udu *UserDeviceUpdate) ClearUpdatedAt() *UserDeviceUpdate {
	udu.mutation.ClearUpdatedAt()
	return udu
}

// SetDeviceID sets the "device_id" field.
func (udu *UserDeviceUpdate) SetDeviceID(s string) *UserDeviceUpdate {
	udu.mutation.SetDeviceID(s)
	return udu
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableDeviceID(s *string) *UserDeviceUpdate {
	if s != nil {
		udu.SetDeviceID(*s)
	}
	return udu
}

// ClearDeviceID clears the value of the "device_id" field.
func (udu *UserDeviceUpdate) ClearDeviceID() *UserDeviceUpdate {
	udu.mutation.ClearDeviceID()
	return udu
}

// Mutation returns the UserDeviceMutation object of the builder.
func (udu *UserDeviceUpdate) Mutation() *UserDeviceMutation {
	return udu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (udu *UserDeviceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(udu.hooks) == 0 {
		affected, err = udu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserDeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			udu.mutation = mutation
			affected, err = udu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(udu.hooks) - 1; i >= 0; i-- {
			if udu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = udu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, udu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (udu *UserDeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := udu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (udu *UserDeviceUpdate) Exec(ctx context.Context) error {
	_, err := udu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udu *UserDeviceUpdate) ExecX(ctx context.Context) {
	if err := udu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (udu *UserDeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userdevice.Table,
			Columns: userdevice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: userdevice.FieldID,
			},
		},
	}
	if ps := udu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := udu.mutation.UserID(); ok {
		_spec.SetField(userdevice.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := udu.mutation.AddedUserID(); ok {
		_spec.AddField(userdevice.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := udu.mutation.Version(); ok {
		_spec.SetField(userdevice.FieldVersion, field.TypeString, value)
	}
	if value, ok := udu.mutation.Platform(); ok {
		_spec.SetField(userdevice.FieldPlatform, field.TypeString, value)
	}
	if value, ok := udu.mutation.LatestSkipUpdate(); ok {
		_spec.SetField(userdevice.FieldLatestSkipUpdate, field.TypeTime, value)
	}
	if udu.mutation.LatestSkipUpdateCleared() {
		_spec.ClearField(userdevice.FieldLatestSkipUpdate, field.TypeTime)
	}
	if value, ok := udu.mutation.CreatedAt(); ok {
		_spec.SetField(userdevice.FieldCreatedAt, field.TypeTime, value)
	}
	if udu.mutation.CreatedAtCleared() {
		_spec.ClearField(userdevice.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := udu.mutation.UpdatedAt(); ok {
		_spec.SetField(userdevice.FieldUpdatedAt, field.TypeTime, value)
	}
	if udu.mutation.UpdatedAtCleared() {
		_spec.ClearField(userdevice.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := udu.mutation.DeviceID(); ok {
		_spec.SetField(userdevice.FieldDeviceID, field.TypeString, value)
	}
	if udu.mutation.DeviceIDCleared() {
		_spec.ClearField(userdevice.FieldDeviceID, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, udu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdevice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserDeviceUpdateOne is the builder for updating a single UserDevice entity.
type UserDeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserDeviceMutation
}

// SetUserID sets the "user_id" field.
func (uduo *UserDeviceUpdateOne) SetUserID(u uint64) *UserDeviceUpdateOne {
	uduo.mutation.ResetUserID()
	uduo.mutation.SetUserID(u)
	return uduo
}

// AddUserID adds u to the "user_id" field.
func (uduo *UserDeviceUpdateOne) AddUserID(u int64) *UserDeviceUpdateOne {
	uduo.mutation.AddUserID(u)
	return uduo
}

// SetVersion sets the "version" field.
func (uduo *UserDeviceUpdateOne) SetVersion(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetVersion(s)
	return uduo
}

// SetPlatform sets the "platform" field.
func (uduo *UserDeviceUpdateOne) SetPlatform(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetPlatform(s)
	return uduo
}

// SetLatestSkipUpdate sets the "latest_skip_update" field.
func (uduo *UserDeviceUpdateOne) SetLatestSkipUpdate(t time.Time) *UserDeviceUpdateOne {
	uduo.mutation.SetLatestSkipUpdate(t)
	return uduo
}

// SetNillableLatestSkipUpdate sets the "latest_skip_update" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableLatestSkipUpdate(t *time.Time) *UserDeviceUpdateOne {
	if t != nil {
		uduo.SetLatestSkipUpdate(*t)
	}
	return uduo
}

// ClearLatestSkipUpdate clears the value of the "latest_skip_update" field.
func (uduo *UserDeviceUpdateOne) ClearLatestSkipUpdate() *UserDeviceUpdateOne {
	uduo.mutation.ClearLatestSkipUpdate()
	return uduo
}

// SetCreatedAt sets the "created_at" field.
func (uduo *UserDeviceUpdateOne) SetCreatedAt(t time.Time) *UserDeviceUpdateOne {
	uduo.mutation.SetCreatedAt(t)
	return uduo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableCreatedAt(t *time.Time) *UserDeviceUpdateOne {
	if t != nil {
		uduo.SetCreatedAt(*t)
	}
	return uduo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (uduo *UserDeviceUpdateOne) ClearCreatedAt() *UserDeviceUpdateOne {
	uduo.mutation.ClearCreatedAt()
	return uduo
}

// SetUpdatedAt sets the "updated_at" field.
func (uduo *UserDeviceUpdateOne) SetUpdatedAt(t time.Time) *UserDeviceUpdateOne {
	uduo.mutation.SetUpdatedAt(t)
	return uduo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserDeviceUpdateOne {
	if t != nil {
		uduo.SetUpdatedAt(*t)
	}
	return uduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uduo *UserDeviceUpdateOne) ClearUpdatedAt() *UserDeviceUpdateOne {
	uduo.mutation.ClearUpdatedAt()
	return uduo
}

// SetDeviceID sets the "device_id" field.
func (uduo *UserDeviceUpdateOne) SetDeviceID(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetDeviceID(s)
	return uduo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableDeviceID(s *string) *UserDeviceUpdateOne {
	if s != nil {
		uduo.SetDeviceID(*s)
	}
	return uduo
}

// ClearDeviceID clears the value of the "device_id" field.
func (uduo *UserDeviceUpdateOne) ClearDeviceID() *UserDeviceUpdateOne {
	uduo.mutation.ClearDeviceID()
	return uduo
}

// Mutation returns the UserDeviceMutation object of the builder.
func (uduo *UserDeviceUpdateOne) Mutation() *UserDeviceMutation {
	return uduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uduo *UserDeviceUpdateOne) Select(field string, fields ...string) *UserDeviceUpdateOne {
	uduo.fields = append([]string{field}, fields...)
	return uduo
}

// Save executes the query and returns the updated UserDevice entity.
func (uduo *UserDeviceUpdateOne) Save(ctx context.Context) (*UserDevice, error) {
	var (
		err  error
		node *UserDevice
	)
	if len(uduo.hooks) == 0 {
		node, err = uduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserDeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uduo.mutation = mutation
			node, err = uduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uduo.hooks) - 1; i >= 0; i-- {
			if uduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserDevice)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserDeviceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uduo *UserDeviceUpdateOne) SaveX(ctx context.Context) *UserDevice {
	node, err := uduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uduo *UserDeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := uduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uduo *UserDeviceUpdateOne) ExecX(ctx context.Context) {
	if err := uduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uduo *UserDeviceUpdateOne) sqlSave(ctx context.Context) (_node *UserDevice, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userdevice.Table,
			Columns: userdevice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: userdevice.FieldID,
			},
		},
	}
	id, ok := uduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserDevice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userdevice.FieldID)
		for _, f := range fields {
			if !userdevice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userdevice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uduo.mutation.UserID(); ok {
		_spec.SetField(userdevice.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := uduo.mutation.AddedUserID(); ok {
		_spec.AddField(userdevice.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := uduo.mutation.Version(); ok {
		_spec.SetField(userdevice.FieldVersion, field.TypeString, value)
	}
	if value, ok := uduo.mutation.Platform(); ok {
		_spec.SetField(userdevice.FieldPlatform, field.TypeString, value)
	}
	if value, ok := uduo.mutation.LatestSkipUpdate(); ok {
		_spec.SetField(userdevice.FieldLatestSkipUpdate, field.TypeTime, value)
	}
	if uduo.mutation.LatestSkipUpdateCleared() {
		_spec.ClearField(userdevice.FieldLatestSkipUpdate, field.TypeTime)
	}
	if value, ok := uduo.mutation.CreatedAt(); ok {
		_spec.SetField(userdevice.FieldCreatedAt, field.TypeTime, value)
	}
	if uduo.mutation.CreatedAtCleared() {
		_spec.ClearField(userdevice.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := uduo.mutation.UpdatedAt(); ok {
		_spec.SetField(userdevice.FieldUpdatedAt, field.TypeTime, value)
	}
	if uduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(userdevice.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := uduo.mutation.DeviceID(); ok {
		_spec.SetField(userdevice.FieldDeviceID, field.TypeString, value)
	}
	if uduo.mutation.DeviceIDCleared() {
		_spec.ClearField(userdevice.FieldDeviceID, field.TypeString)
	}
	_node = &UserDevice{config: uduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdevice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}