// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TDealColumns holds the columns for the "t_deal" table.
	TDealColumns = []*schema.Column{
		{Name: "fdeal_id", Type: field.TypeInt64, Increment: true},
		{Name: "fsku_id", Type: field.TypeInt64, Default: 0},
		{Name: "fproduct_id", Type: field.TypeInt64, Default: 0},
		{Name: "fproduct_name", Type: field.TypeString, Default: ""},
		{Name: "fuid_id", Type: field.TypeInt64, Default: 0},
		{Name: "fcustom_name", Type: field.TypeString, Default: ""},
		{Name: "fphone", Type: field.TypeString, Default: ""},
		{Name: "fprice", Type: field.TypeInt, Default: 0},
		{Name: "faddtime", Type: field.TypeTime},
		{Name: "fmtime", Type: field.TypeTime},
	}
	// TDealTable holds the schema information for the "t_deal" table.
	TDealTable = &schema.Table{
		Name:        "t_deal",
		Columns:     TDealColumns,
		PrimaryKey:  []*schema.Column{TDealColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TDealTable,
	}
)

func init() {
	TDealTable.Annotation = &entsql.Annotation{
		Table: "t_deal",
	}
}
