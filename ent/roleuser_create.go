// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"myapp/ent/roleuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleUserCreate is the builder for creating a RoleUser entity.
type RoleUserCreate struct {
	config
	mutation *RoleUserMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ruc *RoleUserCreate) SetUserID(u uint64) *RoleUserCreate {
	ruc.mutation.SetUserID(u)
	return ruc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ruc *RoleUserCreate) SetNillableUserID(u *uint64) *RoleUserCreate {
	if u != nil {
		ruc.SetUserID(*u)
	}
	return ruc
}

// SetRoleID sets the "role_id" field.
func (ruc *RoleUserCreate) SetRoleID(u uint64) *RoleUserCreate {
	ruc.mutation.SetRoleID(u)
	return ruc
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (ruc *RoleUserCreate) SetNillableRoleID(u *uint64) *RoleUserCreate {
	if u != nil {
		ruc.SetRoleID(*u)
	}
	return ruc
}

// SetCreatedAt sets the "created_at" field.
func (ruc *RoleUserCreate) SetCreatedAt(t time.Time) *RoleUserCreate {
	ruc.mutation.SetCreatedAt(t)
	return ruc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruc *RoleUserCreate) SetNillableCreatedAt(t *time.Time) *RoleUserCreate {
	if t != nil {
		ruc.SetCreatedAt(*t)
	}
	return ruc
}

// SetUpdatedAt sets the "updated_at" field.
func (ruc *RoleUserCreate) SetUpdatedAt(t time.Time) *RoleUserCreate {
	ruc.mutation.SetUpdatedAt(t)
	return ruc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruc *RoleUserCreate) SetNillableUpdatedAt(t *time.Time) *RoleUserCreate {
	if t != nil {
		ruc.SetUpdatedAt(*t)
	}
	return ruc
}

// SetID sets the "id" field.
func (ruc *RoleUserCreate) SetID(u uint64) *RoleUserCreate {
	ruc.mutation.SetID(u)
	return ruc
}

// Mutation returns the RoleUserMutation object of the builder.
func (ruc *RoleUserCreate) Mutation() *RoleUserMutation {
	return ruc.mutation
}

// Save creates the RoleUser in the database.
func (ruc *RoleUserCreate) Save(ctx context.Context) (*RoleUser, error) {
	return withHooks(ctx, ruc.sqlSave, ruc.mutation, ruc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ruc *RoleUserCreate) SaveX(ctx context.Context) *RoleUser {
	v, err := ruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ruc *RoleUserCreate) Exec(ctx context.Context) error {
	_, err := ruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruc *RoleUserCreate) ExecX(ctx context.Context) {
	if err := ruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruc *RoleUserCreate) check() error {
	return nil
}

func (ruc *RoleUserCreate) sqlSave(ctx context.Context) (*RoleUser, error) {
	if err := ruc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ruc.mutation.id = &_node.ID
	ruc.mutation.done = true
	return _node, nil
}

func (ruc *RoleUserCreate) createSpec() (*RoleUser, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleUser{config: ruc.config}
		_spec = sqlgraph.NewCreateSpec(roleuser.Table, sqlgraph.NewFieldSpec(roleuser.FieldID, field.TypeUint64))
	)
	if id, ok := ruc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ruc.mutation.UserID(); ok {
		_spec.SetField(roleuser.FieldUserID, field.TypeUint64, value)
		_node.UserID = value
	}
	if value, ok := ruc.mutation.RoleID(); ok {
		_spec.SetField(roleuser.FieldRoleID, field.TypeUint64, value)
		_node.RoleID = value
	}
	if value, ok := ruc.mutation.CreatedAt(); ok {
		_spec.SetField(roleuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ruc.mutation.UpdatedAt(); ok {
		_spec.SetField(roleuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// RoleUserCreateBulk is the builder for creating many RoleUser entities in bulk.
type RoleUserCreateBulk struct {
	config
	builders []*RoleUserCreate
}

// Save creates the RoleUser entities in the database.
func (rucb *RoleUserCreateBulk) Save(ctx context.Context) ([]*RoleUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rucb.builders))
	nodes := make([]*RoleUser, len(rucb.builders))
	mutators := make([]Mutator, len(rucb.builders))
	for i := range rucb.builders {
		func(i int, root context.Context) {
			builder := rucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rucb *RoleUserCreateBulk) SaveX(ctx context.Context) []*RoleUser {
	v, err := rucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rucb *RoleUserCreateBulk) Exec(ctx context.Context) error {
	_, err := rucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rucb *RoleUserCreateBulk) ExecX(ctx context.Context) {
	if err := rucb.Exec(ctx); err != nil {
		panic(err)
	}
}