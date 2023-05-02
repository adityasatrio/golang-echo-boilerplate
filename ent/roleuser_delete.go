// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"myapp/ent/predicate"
	"myapp/ent/roleuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleUserDelete is the builder for deleting a RoleUser entity.
type RoleUserDelete struct {
	config
	hooks    []Hook
	mutation *RoleUserMutation
}

// Where appends a list predicates to the RoleUserDelete builder.
func (rud *RoleUserDelete) Where(ps ...predicate.RoleUser) *RoleUserDelete {
	rud.mutation.Where(ps...)
	return rud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rud *RoleUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rud.sqlExec, rud.mutation, rud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rud *RoleUserDelete) ExecX(ctx context.Context) int {
	n, err := rud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rud *RoleUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(roleuser.Table, sqlgraph.NewFieldSpec(roleuser.FieldID, field.TypeUint64))
	if ps := rud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rud.mutation.done = true
	return affected, err
}

// RoleUserDeleteOne is the builder for deleting a single RoleUser entity.
type RoleUserDeleteOne struct {
	rud *RoleUserDelete
}

// Where appends a list predicates to the RoleUserDelete builder.
func (rudo *RoleUserDeleteOne) Where(ps ...predicate.RoleUser) *RoleUserDeleteOne {
	rudo.rud.mutation.Where(ps...)
	return rudo
}

// Exec executes the deletion query.
func (rudo *RoleUserDeleteOne) Exec(ctx context.Context) error {
	n, err := rudo.rud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{roleuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rudo *RoleUserDeleteOne) ExecX(ctx context.Context) {
	if err := rudo.Exec(ctx); err != nil {
		panic(err)
	}
}