package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

func (Item) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_item_stock"},
	}
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive().StorageKey("fsku_id"),
		field.Int("totalStock").Default(0).StorageKey("ftotal_stock"),
		field.Int("consumeStock").Default(0).StorageKey("fconsume_Stock"),
		field.Int("leftStock").Default(0).StorageKey("fleft_Stock"),
		field.Time("addtime").Default(time.Now).StorageKey("faddtime"),
		field.Time("mtime").UpdateDefault(time.Now).Default(time.Now).StorageKey("fmtime"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return nil
}
