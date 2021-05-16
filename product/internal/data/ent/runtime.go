// Code generated by entc, DO NOT EDIT.

package ent

import (
	"product/internal/data/ent/product"
	"product/internal/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescAge is the schema descriptor for age field.
	productDescAge := productFields[0].Descriptor()
	// product.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	product.AgeValidator = productDescAge.Validators[0].(func(int) error)
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[1].Descriptor()
	// product.DefaultName holds the default value on creation for the name field.
	product.DefaultName = productDescName.Default.(string)
}
