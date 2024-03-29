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
	hooks     []Hook
	mutation  *RoleUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RoleUserUpdate builder.
func (ruu *RoleUserUpdate) Where(ps ...predicate.RoleUser) *RoleUserUpdate {
	ruu.mutation.Where(ps...)
	return ruu
}

// SetVersions sets the "versions" field.
func (ruu *RoleUserUpdate) SetVersions(i int64) *RoleUserUpdate {
	ruu.mutation.ResetVersions()
	ruu.mutation.SetVersions(i)
	return ruu
}

// SetNillableVersions sets the "versions" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableVersions(i *int64) *RoleUserUpdate {
	if i != nil {
		ruu.SetVersions(*i)
	}
	return ruu
}

// AddVersions adds i to the "versions" field.
func (ruu *RoleUserUpdate) AddVersions(i int64) *RoleUserUpdate {
	ruu.mutation.AddVersions(i)
	return ruu
}

// SetCreatedBy sets the "created_by" field.
func (ruu *RoleUserUpdate) SetCreatedBy(s string) *RoleUserUpdate {
	ruu.mutation.SetCreatedBy(s)
	return ruu
}

// SetUpdatedBy sets the "updated_by" field.
func (ruu *RoleUserUpdate) SetUpdatedBy(s string) *RoleUserUpdate {
	ruu.mutation.SetUpdatedBy(s)
	return ruu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableUpdatedBy(s *string) *RoleUserUpdate {
	if s != nil {
		ruu.SetUpdatedBy(*s)
	}
	return ruu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ruu *RoleUserUpdate) ClearUpdatedBy() *RoleUserUpdate {
	ruu.mutation.ClearUpdatedBy()
	return ruu
}

// SetUpdatedAt sets the "updated_at" field.
func (ruu *RoleUserUpdate) SetUpdatedAt(t time.Time) *RoleUserUpdate {
	ruu.mutation.SetUpdatedAt(t)
	return ruu
}

// SetDeletedBy sets the "deleted_by" field.
func (ruu *RoleUserUpdate) SetDeletedBy(s string) *RoleUserUpdate {
	ruu.mutation.SetDeletedBy(s)
	return ruu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableDeletedBy(s *string) *RoleUserUpdate {
	if s != nil {
		ruu.SetDeletedBy(*s)
	}
	return ruu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ruu *RoleUserUpdate) ClearDeletedBy() *RoleUserUpdate {
	ruu.mutation.ClearDeletedBy()
	return ruu
}

// SetDeletedAt sets the "deleted_at" field.
func (ruu *RoleUserUpdate) SetDeletedAt(t time.Time) *RoleUserUpdate {
	ruu.mutation.SetDeletedAt(t)
	return ruu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableDeletedAt(t *time.Time) *RoleUserUpdate {
	if t != nil {
		ruu.SetDeletedAt(*t)
	}
	return ruu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ruu *RoleUserUpdate) ClearDeletedAt() *RoleUserUpdate {
	ruu.mutation.ClearDeletedAt()
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

// Mutation returns the RoleUserMutation object of the builder.
func (ruu *RoleUserUpdate) Mutation() *RoleUserMutation {
	return ruu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ruu *RoleUserUpdate) Save(ctx context.Context) (int, error) {
	ruu.defaults()
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

// defaults sets the default values of the builder before save.
func (ruu *RoleUserUpdate) defaults() {
	if _, ok := ruu.mutation.UpdatedAt(); !ok {
		v := roleuser.UpdateDefaultUpdatedAt()
		ruu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruu *RoleUserUpdate) check() error {
	if v, ok := ruu.mutation.CreatedBy(); ok {
		if err := roleuser.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "RoleUser.created_by": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruu *RoleUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleUserUpdate {
	ruu.modifiers = append(ruu.modifiers, modifiers...)
	return ruu
}

func (ruu *RoleUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ruu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(roleuser.Table, roleuser.Columns, sqlgraph.NewFieldSpec(roleuser.FieldID, field.TypeUint64))
	if ps := ruu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruu.mutation.Versions(); ok {
		_spec.SetField(roleuser.FieldVersions, field.TypeInt64, value)
	}
	if value, ok := ruu.mutation.AddedVersions(); ok {
		_spec.AddField(roleuser.FieldVersions, field.TypeInt64, value)
	}
	if value, ok := ruu.mutation.CreatedBy(); ok {
		_spec.SetField(roleuser.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := ruu.mutation.UpdatedBy(); ok {
		_spec.SetField(roleuser.FieldUpdatedBy, field.TypeString, value)
	}
	if ruu.mutation.UpdatedByCleared() {
		_spec.ClearField(roleuser.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ruu.mutation.UpdatedAt(); ok {
		_spec.SetField(roleuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruu.mutation.DeletedBy(); ok {
		_spec.SetField(roleuser.FieldDeletedBy, field.TypeString, value)
	}
	if ruu.mutation.DeletedByCleared() {
		_spec.ClearField(roleuser.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ruu.mutation.DeletedAt(); ok {
		_spec.SetField(roleuser.FieldDeletedAt, field.TypeTime, value)
	}
	if ruu.mutation.DeletedAtCleared() {
		_spec.ClearField(roleuser.FieldDeletedAt, field.TypeTime)
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
	_spec.AddModifiers(ruu.modifiers...)
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
	fields    []string
	hooks     []Hook
	mutation  *RoleUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetVersions sets the "versions" field.
func (ruuo *RoleUserUpdateOne) SetVersions(i int64) *RoleUserUpdateOne {
	ruuo.mutation.ResetVersions()
	ruuo.mutation.SetVersions(i)
	return ruuo
}

// SetNillableVersions sets the "versions" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableVersions(i *int64) *RoleUserUpdateOne {
	if i != nil {
		ruuo.SetVersions(*i)
	}
	return ruuo
}

// AddVersions adds i to the "versions" field.
func (ruuo *RoleUserUpdateOne) AddVersions(i int64) *RoleUserUpdateOne {
	ruuo.mutation.AddVersions(i)
	return ruuo
}

// SetCreatedBy sets the "created_by" field.
func (ruuo *RoleUserUpdateOne) SetCreatedBy(s string) *RoleUserUpdateOne {
	ruuo.mutation.SetCreatedBy(s)
	return ruuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruuo *RoleUserUpdateOne) SetUpdatedBy(s string) *RoleUserUpdateOne {
	ruuo.mutation.SetUpdatedBy(s)
	return ruuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableUpdatedBy(s *string) *RoleUserUpdateOne {
	if s != nil {
		ruuo.SetUpdatedBy(*s)
	}
	return ruuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ruuo *RoleUserUpdateOne) ClearUpdatedBy() *RoleUserUpdateOne {
	ruuo.mutation.ClearUpdatedBy()
	return ruuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruuo *RoleUserUpdateOne) SetUpdatedAt(t time.Time) *RoleUserUpdateOne {
	ruuo.mutation.SetUpdatedAt(t)
	return ruuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ruuo *RoleUserUpdateOne) SetDeletedBy(s string) *RoleUserUpdateOne {
	ruuo.mutation.SetDeletedBy(s)
	return ruuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableDeletedBy(s *string) *RoleUserUpdateOne {
	if s != nil {
		ruuo.SetDeletedBy(*s)
	}
	return ruuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ruuo *RoleUserUpdateOne) ClearDeletedBy() *RoleUserUpdateOne {
	ruuo.mutation.ClearDeletedBy()
	return ruuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruuo *RoleUserUpdateOne) SetDeletedAt(t time.Time) *RoleUserUpdateOne {
	ruuo.mutation.SetDeletedAt(t)
	return ruuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableDeletedAt(t *time.Time) *RoleUserUpdateOne {
	if t != nil {
		ruuo.SetDeletedAt(*t)
	}
	return ruuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ruuo *RoleUserUpdateOne) ClearDeletedAt() *RoleUserUpdateOne {
	ruuo.mutation.ClearDeletedAt()
	return ruuo
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
	ruuo.defaults()
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

// defaults sets the default values of the builder before save.
func (ruuo *RoleUserUpdateOne) defaults() {
	if _, ok := ruuo.mutation.UpdatedAt(); !ok {
		v := roleuser.UpdateDefaultUpdatedAt()
		ruuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruuo *RoleUserUpdateOne) check() error {
	if v, ok := ruuo.mutation.CreatedBy(); ok {
		if err := roleuser.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "RoleUser.created_by": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruuo *RoleUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleUserUpdateOne {
	ruuo.modifiers = append(ruuo.modifiers, modifiers...)
	return ruuo
}

func (ruuo *RoleUserUpdateOne) sqlSave(ctx context.Context) (_node *RoleUser, err error) {
	if err := ruuo.check(); err != nil {
		return _node, err
	}
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
	if value, ok := ruuo.mutation.Versions(); ok {
		_spec.SetField(roleuser.FieldVersions, field.TypeInt64, value)
	}
	if value, ok := ruuo.mutation.AddedVersions(); ok {
		_spec.AddField(roleuser.FieldVersions, field.TypeInt64, value)
	}
	if value, ok := ruuo.mutation.CreatedBy(); ok {
		_spec.SetField(roleuser.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := ruuo.mutation.UpdatedBy(); ok {
		_spec.SetField(roleuser.FieldUpdatedBy, field.TypeString, value)
	}
	if ruuo.mutation.UpdatedByCleared() {
		_spec.ClearField(roleuser.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ruuo.mutation.UpdatedAt(); ok {
		_spec.SetField(roleuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruuo.mutation.DeletedBy(); ok {
		_spec.SetField(roleuser.FieldDeletedBy, field.TypeString, value)
	}
	if ruuo.mutation.DeletedByCleared() {
		_spec.ClearField(roleuser.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ruuo.mutation.DeletedAt(); ok {
		_spec.SetField(roleuser.FieldDeletedAt, field.TypeTime, value)
	}
	if ruuo.mutation.DeletedAtCleared() {
		_spec.ClearField(roleuser.FieldDeletedAt, field.TypeTime)
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
	_spec.AddModifiers(ruuo.modifiers...)
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
