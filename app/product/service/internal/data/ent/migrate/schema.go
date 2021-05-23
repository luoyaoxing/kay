// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TProductColumns holds the columns for the "t_product" table.
	TProductColumns = []*schema.Column{
		{Name: "fsku_id", Type: field.TypeInt64, Increment: true},
		{Name: "fversion_id", Type: field.TypeInt64, Default: 0},
		{Name: "fversion_name", Type: field.TypeString, Default: ""},
		{Name: "fproduct_id", Type: field.TypeInt64, Default: 0},
		{Name: "fproduct_name", Type: field.TypeString, Default: ""},
		{Name: "fprice", Type: field.TypeInt, Default: 0},
		{Name: "fattr", Type: field.TypeString, Default: ""},
		{Name: "fproduct_desc", Type: field.TypeString, Default: ""},
		{Name: "faddtime", Type: field.TypeTime},
		{Name: "fmtime", Type: field.TypeTime},
	}
	// TProductTable holds the schema information for the "t_product" table.
	TProductTable = &schema.Table{
		Name:        "t_product",
		Columns:     TProductColumns,
		PrimaryKey:  []*schema.Column{TProductColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TProductTable,
	}
)

func init() {
	TProductTable.Annotation = &entsql.Annotation{
		Table: "t_product",
	}
}
