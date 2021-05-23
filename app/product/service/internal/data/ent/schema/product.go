package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_product"},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive().StorageKey("fsku_id"),
		field.Int64("versionId").Positive().Default(0).StorageKey("fversion_id"),
		field.String("versionName").NotEmpty().Default("").StorageKey("fversion_name"),
		field.Int64("productId").Positive().Default(0).StorageKey("fproduct_id"),
		field.String("productName").NotEmpty().Default("").StorageKey("fproduct_name"),
		field.Int("price").Positive().Default(0).StorageKey("fprice"),
		field.String("attr").NotEmpty().Default("").StorageKey("fattr"),
		field.String("productDesc").NotEmpty().Default("").StorageKey("fproduct_desc"),
		field.Time("addtime").Default(time.Now).StorageKey("faddtime"),
		field.Time("mtime").UpdateDefault(time.Now).Default(time.Now).StorageKey("fmtime"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
