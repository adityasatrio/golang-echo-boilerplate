// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myapp/ent/systemparameter"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SystemParameterCreate is the builder for creating a SystemParameter entity.
type SystemParameterCreate struct {
	config
	mutation *SystemParameterMutation
	hooks    []Hook
}

// SetKey sets the "key" field.
func (spc *SystemParameterCreate) SetKey(s string) *SystemParameterCreate {
	spc.mutation.SetKey(s)
	return spc
}

// SetValue sets the "value" field.
func (spc *SystemParameterCreate) SetValue(s string) *SystemParameterCreate {
	spc.mutation.SetValue(s)
	return spc
}

// SetCreatedBy sets the "created_by" field.
func (spc *SystemParameterCreate) SetCreatedBy(s string) *SystemParameterCreate {
	spc.mutation.SetCreatedBy(s)
	return spc
}

// SetCreatedAt sets the "created_at" field.
func (spc *SystemParameterCreate) SetCreatedAt(t time.Time) *SystemParameterCreate {
	spc.mutation.SetCreatedAt(t)
	return spc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (spc *SystemParameterCreate) SetNillableCreatedAt(t *time.Time) *SystemParameterCreate {
	if t != nil {
		spc.SetCreatedAt(*t)
	}
	return spc
}

// SetUpdatedBy sets the "updated_by" field.
func (spc *SystemParameterCreate) SetUpdatedBy(s string) *SystemParameterCreate {
	spc.mutation.SetUpdatedBy(s)
	return spc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (spc *SystemParameterCreate) SetNillableUpdatedBy(s *string) *SystemParameterCreate {
	if s != nil {
		spc.SetUpdatedBy(*s)
	}
	return spc
}

// SetUpdatedAt sets the "updated_at" field.
func (spc *SystemParameterCreate) SetUpdatedAt(t time.Time) *SystemParameterCreate {
	spc.mutation.SetUpdatedAt(t)
	return spc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (spc *SystemParameterCreate) SetNillableUpdatedAt(t *time.Time) *SystemParameterCreate {
	if t != nil {
		spc.SetUpdatedAt(*t)
	}
	return spc
}

// SetDeletedBy sets the "deleted_by" field.
func (spc *SystemParameterCreate) SetDeletedBy(s string) *SystemParameterCreate {
	spc.mutation.SetDeletedBy(s)
	return spc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (spc *SystemParameterCreate) SetNillableDeletedBy(s *string) *SystemParameterCreate {
	if s != nil {
		spc.SetDeletedBy(*s)
	}
	return spc
}

// SetDeletedAt sets the "deleted_at" field.
func (spc *SystemParameterCreate) SetDeletedAt(t time.Time) *SystemParameterCreate {
	spc.mutation.SetDeletedAt(t)
	return spc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (spc *SystemParameterCreate) SetNillableDeletedAt(t *time.Time) *SystemParameterCreate {
	if t != nil {
		spc.SetDeletedAt(*t)
	}
	return spc
}

// Mutation returns the SystemParameterMutation object of the builder.
func (spc *SystemParameterCreate) Mutation() *SystemParameterMutation {
	return spc.mutation
}

// Save creates the SystemParameter in the database.
func (spc *SystemParameterCreate) Save(ctx context.Context) (*SystemParameter, error) {
	spc.defaults()
	return withHooks(ctx, spc.sqlSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *SystemParameterCreate) SaveX(ctx context.Context) *SystemParameter {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *SystemParameterCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *SystemParameterCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spc *SystemParameterCreate) defaults() {
	if _, ok := spc.mutation.CreatedAt(); !ok {
		v := systemparameter.DefaultCreatedAt
		spc.mutation.SetCreatedAt(v)
	}
	if _, ok := spc.mutation.UpdatedAt(); !ok {
		v := systemparameter.DefaultUpdatedAt
		spc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *SystemParameterCreate) check() error {
	if _, ok := spc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "SystemParameter.key"`)}
	}
	if v, ok := spc.mutation.Key(); ok {
		if err := systemparameter.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "SystemParameter.key": %w`, err)}
		}
	}
	if _, ok := spc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "SystemParameter.value"`)}
	}
	if v, ok := spc.mutation.Value(); ok {
		if err := systemparameter.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "SystemParameter.value": %w`, err)}
		}
	}
	if _, ok := spc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "SystemParameter.created_by"`)}
	}
	if v, ok := spc.mutation.CreatedBy(); ok {
		if err := systemparameter.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "SystemParameter.created_by": %w`, err)}
		}
	}
	if _, ok := spc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SystemParameter.created_at"`)}
	}
	return nil
}

func (spc *SystemParameterCreate) sqlSave(ctx context.Context) (*SystemParameter, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	spc.mutation.id = &_node.ID
	spc.mutation.done = true
	return _node, nil
}

func (spc *SystemParameterCreate) createSpec() (*SystemParameter, *sqlgraph.CreateSpec) {
	var (
		_node = &SystemParameter{config: spc.config}
		_spec = sqlgraph.NewCreateSpec(systemparameter.Table, sqlgraph.NewFieldSpec(systemparameter.FieldID, field.TypeInt))
	)
	if value, ok := spc.mutation.Key(); ok {
		_spec.SetField(systemparameter.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := spc.mutation.Value(); ok {
		_spec.SetField(systemparameter.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := spc.mutation.CreatedBy(); ok {
		_spec.SetField(systemparameter.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := spc.mutation.CreatedAt(); ok {
		_spec.SetField(systemparameter.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := spc.mutation.UpdatedBy(); ok {
		_spec.SetField(systemparameter.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := spc.mutation.UpdatedAt(); ok {
		_spec.SetField(systemparameter.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := spc.mutation.DeletedBy(); ok {
		_spec.SetField(systemparameter.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := spc.mutation.DeletedAt(); ok {
		_spec.SetField(systemparameter.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	return _node, _spec
}

// SystemParameterCreateBulk is the builder for creating many SystemParameter entities in bulk.
type SystemParameterCreateBulk struct {
	config
	builders []*SystemParameterCreate
}

// Save creates the SystemParameter entities in the database.
func (spcb *SystemParameterCreateBulk) Save(ctx context.Context) ([]*SystemParameter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*SystemParameter, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SystemParameterMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *SystemParameterCreateBulk) SaveX(ctx context.Context) []*SystemParameter {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *SystemParameterCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *SystemParameterCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}
