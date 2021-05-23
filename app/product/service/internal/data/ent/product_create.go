// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kay/app/product/service/internal/data/ent/product"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetVersionId sets the "versionId" field.
func (pc *ProductCreate) SetVersionId(i int64) *ProductCreate {
	pc.mutation.SetVersionId(i)
	return pc
}

// SetNillableVersionId sets the "versionId" field if the given value is not nil.
func (pc *ProductCreate) SetNillableVersionId(i *int64) *ProductCreate {
	if i != nil {
		pc.SetVersionId(*i)
	}
	return pc
}

// SetVersionName sets the "versionName" field.
func (pc *ProductCreate) SetVersionName(s string) *ProductCreate {
	pc.mutation.SetVersionName(s)
	return pc
}

// SetNillableVersionName sets the "versionName" field if the given value is not nil.
func (pc *ProductCreate) SetNillableVersionName(s *string) *ProductCreate {
	if s != nil {
		pc.SetVersionName(*s)
	}
	return pc
}

// SetProductId sets the "productId" field.
func (pc *ProductCreate) SetProductId(i int64) *ProductCreate {
	pc.mutation.SetProductId(i)
	return pc
}

// SetNillableProductId sets the "productId" field if the given value is not nil.
func (pc *ProductCreate) SetNillableProductId(i *int64) *ProductCreate {
	if i != nil {
		pc.SetProductId(*i)
	}
	return pc
}

// SetProductName sets the "productName" field.
func (pc *ProductCreate) SetProductName(s string) *ProductCreate {
	pc.mutation.SetProductName(s)
	return pc
}

// SetNillableProductName sets the "productName" field if the given value is not nil.
func (pc *ProductCreate) SetNillableProductName(s *string) *ProductCreate {
	if s != nil {
		pc.SetProductName(*s)
	}
	return pc
}

// SetPrice sets the "price" field.
func (pc *ProductCreate) SetPrice(i int) *ProductCreate {
	pc.mutation.SetPrice(i)
	return pc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pc *ProductCreate) SetNillablePrice(i *int) *ProductCreate {
	if i != nil {
		pc.SetPrice(*i)
	}
	return pc
}

// SetAttr sets the "attr" field.
func (pc *ProductCreate) SetAttr(s string) *ProductCreate {
	pc.mutation.SetAttr(s)
	return pc
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (pc *ProductCreate) SetNillableAttr(s *string) *ProductCreate {
	if s != nil {
		pc.SetAttr(*s)
	}
	return pc
}

// SetProductDesc sets the "productDesc" field.
func (pc *ProductCreate) SetProductDesc(s string) *ProductCreate {
	pc.mutation.SetProductDesc(s)
	return pc
}

// SetNillableProductDesc sets the "productDesc" field if the given value is not nil.
func (pc *ProductCreate) SetNillableProductDesc(s *string) *ProductCreate {
	if s != nil {
		pc.SetProductDesc(*s)
	}
	return pc
}

// SetAddtime sets the "addtime" field.
func (pc *ProductCreate) SetAddtime(t time.Time) *ProductCreate {
	pc.mutation.SetAddtime(t)
	return pc
}

// SetNillableAddtime sets the "addtime" field if the given value is not nil.
func (pc *ProductCreate) SetNillableAddtime(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetAddtime(*t)
	}
	return pc
}

// SetMtime sets the "mtime" field.
func (pc *ProductCreate) SetMtime(t time.Time) *ProductCreate {
	pc.mutation.SetMtime(t)
	return pc
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (pc *ProductCreate) SetNillableMtime(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetMtime(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProductCreate) SetID(i int64) *ProductCreate {
	pc.mutation.SetID(i)
	return pc
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.VersionId(); !ok {
		v := product.DefaultVersionId
		pc.mutation.SetVersionId(v)
	}
	if _, ok := pc.mutation.VersionName(); !ok {
		v := product.DefaultVersionName
		pc.mutation.SetVersionName(v)
	}
	if _, ok := pc.mutation.ProductId(); !ok {
		v := product.DefaultProductId
		pc.mutation.SetProductId(v)
	}
	if _, ok := pc.mutation.ProductName(); !ok {
		v := product.DefaultProductName
		pc.mutation.SetProductName(v)
	}
	if _, ok := pc.mutation.Price(); !ok {
		v := product.DefaultPrice
		pc.mutation.SetPrice(v)
	}
	if _, ok := pc.mutation.Attr(); !ok {
		v := product.DefaultAttr
		pc.mutation.SetAttr(v)
	}
	if _, ok := pc.mutation.ProductDesc(); !ok {
		v := product.DefaultProductDesc
		pc.mutation.SetProductDesc(v)
	}
	if _, ok := pc.mutation.Addtime(); !ok {
		v := product.DefaultAddtime()
		pc.mutation.SetAddtime(v)
	}
	if _, ok := pc.mutation.Mtime(); !ok {
		v := product.DefaultMtime()
		pc.mutation.SetMtime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.VersionId(); !ok {
		return &ValidationError{Name: "versionId", err: errors.New("ent: missing required field \"versionId\"")}
	}
	if v, ok := pc.mutation.VersionId(); ok {
		if err := product.VersionIdValidator(v); err != nil {
			return &ValidationError{Name: "versionId", err: fmt.Errorf("ent: validator failed for field \"versionId\": %w", err)}
		}
	}
	if _, ok := pc.mutation.VersionName(); !ok {
		return &ValidationError{Name: "versionName", err: errors.New("ent: missing required field \"versionName\"")}
	}
	if v, ok := pc.mutation.VersionName(); ok {
		if err := product.VersionNameValidator(v); err != nil {
			return &ValidationError{Name: "versionName", err: fmt.Errorf("ent: validator failed for field \"versionName\": %w", err)}
		}
	}
	if _, ok := pc.mutation.ProductId(); !ok {
		return &ValidationError{Name: "productId", err: errors.New("ent: missing required field \"productId\"")}
	}
	if v, ok := pc.mutation.ProductId(); ok {
		if err := product.ProductIdValidator(v); err != nil {
			return &ValidationError{Name: "productId", err: fmt.Errorf("ent: validator failed for field \"productId\": %w", err)}
		}
	}
	if _, ok := pc.mutation.ProductName(); !ok {
		return &ValidationError{Name: "productName", err: errors.New("ent: missing required field \"productName\"")}
	}
	if v, ok := pc.mutation.ProductName(); ok {
		if err := product.ProductNameValidator(v); err != nil {
			return &ValidationError{Name: "productName", err: fmt.Errorf("ent: validator failed for field \"productName\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if v, ok := pc.mutation.Price(); ok {
		if err := product.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf("ent: validator failed for field \"price\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Attr(); !ok {
		return &ValidationError{Name: "attr", err: errors.New("ent: missing required field \"attr\"")}
	}
	if v, ok := pc.mutation.Attr(); ok {
		if err := product.AttrValidator(v); err != nil {
			return &ValidationError{Name: "attr", err: fmt.Errorf("ent: validator failed for field \"attr\": %w", err)}
		}
	}
	if _, ok := pc.mutation.ProductDesc(); !ok {
		return &ValidationError{Name: "productDesc", err: errors.New("ent: missing required field \"productDesc\"")}
	}
	if v, ok := pc.mutation.ProductDesc(); ok {
		if err := product.ProductDescValidator(v); err != nil {
			return &ValidationError{Name: "productDesc", err: fmt.Errorf("ent: validator failed for field \"productDesc\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Addtime(); !ok {
		return &ValidationError{Name: "addtime", err: errors.New("ent: missing required field \"addtime\"")}
	}
	if _, ok := pc.mutation.Mtime(); !ok {
		return &ValidationError{Name: "mtime", err: errors.New("ent: missing required field \"mtime\"")}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := product.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: product.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.VersionId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: product.FieldVersionId,
		})
		_node.VersionId = value
	}
	if value, ok := pc.mutation.VersionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldVersionName,
		})
		_node.VersionName = value
	}
	if value, ok := pc.mutation.ProductId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: product.FieldProductId,
		})
		_node.ProductId = value
	}
	if value, ok := pc.mutation.ProductName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldProductName,
		})
		_node.ProductName = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := pc.mutation.Attr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldAttr,
		})
		_node.Attr = value
	}
	if value, ok := pc.mutation.ProductDesc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldProductDesc,
		})
		_node.ProductDesc = value
	}
	if value, ok := pc.mutation.Addtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldAddtime,
		})
		_node.Addtime = value
	}
	if value, ok := pc.mutation.Mtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldMtime,
		})
		_node.Mtime = value
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
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
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}