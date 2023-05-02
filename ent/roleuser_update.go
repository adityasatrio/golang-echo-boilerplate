// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myapp/ent/predicate"
	"myapp/ent/roleuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleUserUpdate is the builder for updating RoleUser entities.
type RoleUserUpdate struct {
	config
	hooks    []Hook
	mutation *RoleUserMutation
}

// Where appends a list predicates to the RoleUserUpdate builder.
func (ruu *RoleUserUpdate) Where(ps ...predicate.RoleUser) *RoleUserUpdate {
	ruu.mutation.Where(ps...)
	return ruu
}

// SetUserID sets the "user_id" field.
func (ruu *RoleUserUpdate) SetUserID(u uint64) *RoleUserUpdate {
	ruu.mutation.ResetUserID()
	ruu.mutation.SetUserID(u)
	return ruu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableUserID(u *uint64) *RoleUserUpdate {
	if u != nil {
		ruu.SetUserID(*u)
	}
	return ruu
}

// AddUserID adds u to the "user_id" field.
func (ruu *RoleUserUpdate) AddUserID(u int64) *RoleUserUpdate {
	ruu.mutation.AddUserID(u)
	return ruu
}

// ClearUserID clears the value of the "user_id" field.
func (ruu *RoleUserUpdate) ClearUserID() *RoleUserUpdate {
	ruu.mutation.ClearUserID()
	return ruu
}

// SetRoleID sets the "role_id" field.
func (ruu *RoleUserUpdate) SetRoleID(u uint64) *RoleUserUpdate {
	ruu.mutation.ResetRoleID()
	ruu.mutation.SetRoleID(u)
	return ruu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableRoleID(u *uint64) *RoleUserUpdate {
	if u != nil {
		ruu.SetRoleID(*u)
	}
	return ruu
}

// AddRoleID adds u to the "role_id" field.
func (ruu *RoleUserUpdate) AddRoleID(u int64) *RoleUserUpdate {
	ruu.mutation.AddRoleID(u)
	return ruu
}

// ClearRoleID clears the value of the "role_id" field.
func (ruu *RoleUserUpdate) ClearRoleID() *RoleUserUpdate {
	ruu.mutation.ClearRoleID()
	return ruu
}

// SetCreatedAt sets the "created_at" field.
func (ruu *RoleUserUpdate) SetCreatedAt(t time.Time) *RoleUserUpdate {
	ruu.mutation.SetCreatedAt(t)
	return ruu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableCreatedAt(t *time.Time) *RoleUserUpdate {
	if t != nil {
		ruu.SetCreatedAt(*t)
	}
	return ruu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (ruu *RoleUserUpdate) ClearCreatedAt() *RoleUserUpdate {
	ruu.mutation.ClearCreatedAt()
	return ruu
}

// SetUpdatedAt sets the "updated_at" field.
func (ruu *RoleUserUpdate) SetUpdatedAt(t time.Time) *RoleUserUpdate {
	ruu.mutation.SetUpdatedAt(t)
	return ruu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableUpdatedAt(t *time.Time) *RoleUserUpdate {
	if t != nil {
		ruu.SetUpdatedAt(*t)
	}
	return ruu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ruu *RoleUserUpdate) ClearUpdatedAt() *RoleUserUpdate {
	ruu.mutation.ClearUpdatedAt()
	return ruu
}

// Mutation returns the RoleUserMutation object of the builder.
func (ruu *RoleUserUpdate) Mutation() *RoleUserMutation {
	return ruu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ruu *RoleUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ruu.sqlSave, ruu.mutation, ruu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruu *RoleUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ruu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ruu *RoleUserUpdate) Exec(ctx context.Context) error {
	_, err := ruu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruu *RoleUserUpdate) ExecX(ctx context.Context) {
	if err := ruu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruu *RoleUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(roleuser.Table, roleuser.Columns, sqlgraph.NewFieldSpec(roleuser.FieldID, field.TypeUint64))
	if ps := ruu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruu.mutation.UserID(); ok {
		_spec.SetField(roleuser.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := ruu.mutation.AddedUserID(); ok {
		_spec.AddField(roleuser.FieldUserID, field.TypeUint64, value)
	}
	if ruu.mutation.UserIDCleared() {
		_spec.ClearField(roleuser.FieldUserID, field.TypeUint64)
	}
	if value, ok := ruu.mutation.RoleID(); ok {
		_spec.SetField(roleuser.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := ruu.mutation.AddedRoleID(); ok {
		_spec.AddField(roleuser.FieldRoleID, field.TypeUint64, value)
	}
	if ruu.mutation.RoleIDCleared() {
		_spec.ClearField(roleuser.FieldRoleID, field.TypeUint64)
	}
	if value, ok := ruu.mutation.CreatedAt(); ok {
		_spec.SetField(roleuser.FieldCreatedAt, field.TypeTime, value)
	}
	if ruu.mutation.CreatedAtCleared() {
		_spec.ClearField(roleuser.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ruu.mutation.UpdatedAt(); ok {
		_spec.SetField(roleuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if ruu.mutation.UpdatedAtCleared() {
		_spec.ClearField(roleuser.FieldUpdatedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ruu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ruu.mutation.done = true
	return n, nil
}

// RoleUserUpdateOne is the builder for updating a single RoleUser entity.
type RoleUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleUserMutation
}

// SetUserID sets the "user_id" field.
func (ruuo *RoleUserUpdateOne) SetUserID(u uint64) *RoleUserUpdateOne {
	ruuo.mutation.ResetUserID()
	ruuo.mutation.SetUserID(u)
	return ruuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableUserID(u *uint64) *RoleUserUpdateOne {
	if u != nil {
		ruuo.SetUserID(*u)
	}
	return ruuo
}

// AddUserID adds u to the "user_id" field.
func (ruuo *RoleUserUpdateOne) AddUserID(u int64) *RoleUserUpdateOne {
	ruuo.mutation.AddUserID(u)
	return ruuo
}

// ClearUserID clears the value of the "user_id" field.
func (ruuo *RoleUserUpdateOne) ClearUserID() *RoleUserUpdateOne {
	ruuo.mutation.ClearUserID()
	return ruuo
}

// SetRoleID sets the "role_id" field.
func (ruuo *RoleUserUpdateOne) SetRoleID(u uint64) *RoleUserUpdateOne {
	ruuo.mutation.ResetRoleID()
	ruuo.mutation.SetRoleID(u)
	return ruuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableRoleID(u *uint64) *RoleUserUpdateOne {
	if u != nil {
		ruuo.SetRoleID(*u)
	}
	return ruuo
}

// AddRoleID adds u to the "role_id" field.
func (ruuo *RoleUserUpdateOne) AddRoleID(u int64) *RoleUserUpdateOne {
	ruuo.mutation.AddRoleID(u)
	return ruuo
}

// ClearRoleID clears the value of the "role_id" field.
func (ruuo *RoleUserUpdateOne) ClearRoleID() *RoleUserUpdateOne {
	ruuo.mutation.ClearRoleID()
	return ruuo
}

// SetCreatedAt sets the "created_at" field.
func (ruuo *RoleUserUpdateOne) SetCreatedAt(t time.Time) *RoleUserUpdateOne {
	ruuo.mutation.SetCreatedAt(t)
	return ruuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableCreatedAt(t *time.Time) *RoleUserUpdateOne {
	if t != nil {
		ruuo.SetCreatedAt(*t)
	}
	return ruuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (ruuo *RoleUserUpdateOne) ClearCreatedAt() *RoleUserUpdateOne {
	ruuo.mutation.ClearCreatedAt()
	return ruuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruuo *RoleUserUpdateOne) SetUpdatedAt(t time.Time) *RoleUserUpdateOne {
	ruuo.mutation.SetUpdatedAt(t)
	return ruuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableUpdatedAt(t *time.Time) *RoleUserUpdateOne {
	if t != nil {
		ruuo.SetUpdatedAt(*t)
	}
	return ruuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ruuo *RoleUserUpdateOne) ClearUpdatedAt() *RoleUserUpdateOne {
	ruuo.mutation.ClearUpdatedAt()
	return ruuo
}

// Mutation returns the RoleUserMutation object of the builder.
func (ruuo *RoleUserUpdateOne) Mutation() *RoleUserMutation {
	return ruuo.mutation
}

// Where appends a list predicates to the RoleUserUpdate builder.
func (ruuo *RoleUserUpdateOne) Where(ps ...predicate.RoleUser) *RoleUserUpdateOne {
	ruuo.mutation.Where(ps...)
	return ruuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruuo *RoleUserUpdateOne) Select(field string, fields ...string) *RoleUserUpdateOne {
	ruuo.fields = append([]string{field}, fields...)
	return ruuo
}

// Save executes the query and returns the updated RoleUser entity.
func (ruuo *RoleUserUpdateOne) Save(ctx context.Context) (*RoleUser, error) {
	return withHooks(ctx, ruuo.sqlSave, ruuo.mutation, ruuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruuo *RoleUserUpdateOne) SaveX(ctx context.Context) *RoleUser {
	node, err := ruuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruuo *RoleUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ruuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruuo *RoleUserUpdateOne) ExecX(ctx context.Context) {
	if err := ruuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruuo *RoleUserUpdateOne) sqlSave(ctx context.Context) (_node *RoleUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(roleuser.Table, roleuser.Columns, sqlgraph.NewFieldSpec(roleuser.FieldID, field.TypeUint64))
	id, ok := ruuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoleUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roleuser.FieldID)
		for _, f := range fields {
			if !roleuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != roleuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruuo.mutation.UserID(); ok {
		_spec.SetField(roleuser.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := ruuo.mutation.AddedUserID(); ok {
		_spec.AddField(roleuser.FieldUserID, field.TypeUint64, value)
	}
	if ruuo.mutation.UserIDCleared() {
		_spec.ClearField(roleuser.FieldUserID, field.TypeUint64)
	}
	if value, ok := ruuo.mutation.RoleID(); ok {
		_spec.SetField(roleuser.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := ruuo.mutation.AddedRoleID(); ok {
		_spec.AddField(roleuser.FieldRoleID, field.TypeUint64, value)
	}
	if ruuo.mutation.RoleIDCleared() {
		_spec.ClearField(roleuser.FieldRoleID, field.TypeUint64)
	}
	if value, ok := ruuo.mutation.CreatedAt(); ok {
		_spec.SetField(roleuser.FieldCreatedAt, field.TypeTime, value)
	}
	if ruuo.mutation.CreatedAtCleared() {
		_spec.ClearField(roleuser.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ruuo.mutation.UpdatedAt(); ok {
		_spec.SetField(roleuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if ruuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(roleuser.FieldUpdatedAt, field.TypeTime)
	}
	_node = &RoleUser{config: ruuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruuo.mutation.done = true
	return _node, nil
}