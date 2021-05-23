package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Deal holds the schema definition for the Deal entity.
type Deal struct {
	ent.Schema
}

func (Deal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_deal"},
	}
}

// Fields of the Deal.
func (Deal) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive().StorageKey("fdeal_id"),
		field.Int64("skuId").Positive().Default(0).StorageKey("fsku_id"),
		field.Int64("productId").Positive().Default(0).StorageKey("fproduct_id"),
		field.String("productName").NotEmpty().Default("").StorageKey("fproduct_name"),
		field.Int64("uid").Positive().Default(0).StorageKey("fuid_id"),
		field.String("customName").NotEmpty().Default("").StorageKey("fcustom_name"),
		field.String("phone").NotEmpty().Default("").StorageKey("fphone"),
		field.Int("price").Positive().Default(0).StorageKey("fprice"),
		field.Time("addtime").Default(time.Now).StorageKey("faddtime"),
		field.Time("mtime").UpdateDefault(time.Now).Default(time.Now).StorageKey("fmtime"),
	}
}

// Edges of the Deal.
func (Deal) Edges() []ent.Edge {
	return nil
}
