// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myapp/ent/pet"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (pc *PetCreate) SetVersion(i int64) *PetCreate {
	pc.mutation.SetVersion(i)
	return pc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (pc *PetCreate) SetNillableVersion(i *int64) *PetCreate {
	if i != nil {
		pc.SetVersion(*i)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *PetCreate) SetCreatedBy(s string) *PetCreate {
	pc.mutation.SetCreatedBy(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PetCreate) SetCreatedAt(t time.Time) *PetCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PetCreate) SetNillableCreatedAt(t *time.Time) *PetCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PetCreate) SetUpdatedBy(s string) *PetCreate {
	pc.mutation.SetUpdatedBy(s)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PetCreate) SetNillableUpdatedBy(s *string) *PetCreate {
	if s != nil {
		pc.SetUpdatedBy(*s)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PetCreate) SetUpdatedAt(t time.Time) *PetCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PetCreate) SetNillableUpdatedAt(t *time.Time) *PetCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedBy sets the "deleted_by" field.
func (pc *PetCreate) SetDeletedBy(s string) *PetCreate {
	pc.mutation.SetDeletedBy(s)
	return pc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (pc *PetCreate) SetNillableDeletedBy(s *string) *PetCreate {
	if s != nil {
		pc.SetDeletedBy(*s)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PetCreate) SetDeletedAt(t time.Time) *PetCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PetCreate) SetNillableDeletedAt(t *time.Time) *PetCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetType sets the "type" field.
func (pc *PetCreate) SetType(pe pet.Type) *PetCreate {
	pc.mutation.SetType(pe)
	return pc
}

// SetCode sets the "code" field.
func (pc *PetCreate) SetCode(s string) *PetCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetAgeMonth sets the "age_month" field.
func (pc *PetCreate) SetAgeMonth(i int) *PetCreate {
	pc.mutation.SetAgeMonth(i)
	return pc
}

// SetID sets the "id" field.
func (pc *PetCreate) SetID(u uuid.UUID) *PetCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PetCreate) SetNillableID(u *uuid.UUID) *PetCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// Mutation returns the PetMutation object of the builder.
func (pc *PetCreate) Mutation() *PetMutation {
	return pc.mutation
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PetCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PetCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PetCreate) defaults() {
	if _, ok := pc.mutation.Version(); !ok {
		v := pet.DefaultVersion()
		pc.mutation.SetVersion(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := pet.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := pet.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := pet.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PetCreate) check() error {
	if _, ok := pc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Pet.version"`)}
	}
	if _, ok := pc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Pet.created_by"`)}
	}
	if v, ok := pc.mutation.CreatedBy(); ok {
		if err := pet.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "Pet.created_by": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Pet.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Pet.updated_at"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := pet.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Pet.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Pet.type"`)}
	}
	if v, ok := pc.mutation.GetType(); ok {
		if err := pet.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Pet.type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Pet.code"`)}
	}
	if v, ok := pc.mutation.Code(); ok {
		if err := pet.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Pet.code": %w`, err)}
		}
	}
	if _, ok := pc.mutation.AgeMonth(); !ok {
		return &ValidationError{Name: "age_month", err: errors.New(`ent: missing required field "Pet.age_month"`)}
	}
	if v, ok := pc.mutation.AgeMonth(); ok {
		if err := pet.AgeMonthValidator(v); err != nil {
			return &ValidationError{Name: "age_month", err: fmt.Errorf(`ent: validator failed for field "Pet.age_month": %w`, err)}
		}
	}
	return nil
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(pet.Table, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Version(); ok {
		_spec.SetField(pet.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(pet.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(pet.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(pet.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(pet.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedBy(); ok {
		_spec.SetField(pet.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(pet.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.SetField(pet.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(pet.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := pc.mutation.AgeMonth(); ok {
		_spec.SetField(pet.FieldAgeMonth, field.TypeInt, value)
		_node.AgeMonth = value
	}
	return _node, _spec
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	builders []*PetCreate
}

// Save creates the Pet entities in the database.
func (pcb *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pet, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PetCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PetCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
