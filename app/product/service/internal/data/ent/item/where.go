// Code generated by entc, DO NOT EDIT.

package item

import (
	"kay/app/product/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TotalStock applies equality check predicate on the "totalStock" field. It's identical to TotalStockEQ.
func TotalStock(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalStock), v))
	})
}

// ConsumeStock applies equality check predicate on the "consumeStock" field. It's identical to ConsumeStockEQ.
func ConsumeStock(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConsumeStock), v))
	})
}

// LeftStock applies equality check predicate on the "leftStock" field. It's identical to LeftStockEQ.
func LeftStock(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeftStock), v))
	})
}

// Addtime applies equality check predicate on the "addtime" field. It's identical to AddtimeEQ.
func Addtime(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddtime), v))
	})
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// TotalStockEQ applies the EQ predicate on the "totalStock" field.
func TotalStockEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalStock), v))
	})
}

// TotalStockNEQ applies the NEQ predicate on the "totalStock" field.
func TotalStockNEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalStock), v))
	})
}

// TotalStockIn applies the In predicate on the "totalStock" field.
func TotalStockIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTotalStock), v...))
	})
}

// TotalStockNotIn applies the NotIn predicate on the "totalStock" field.
func TotalStockNotIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTotalStock), v...))
	})
}

// TotalStockGT applies the GT predicate on the "totalStock" field.
func TotalStockGT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalStock), v))
	})
}

// TotalStockGTE applies the GTE predicate on the "totalStock" field.
func TotalStockGTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalStock), v))
	})
}

// TotalStockLT applies the LT predicate on the "totalStock" field.
func TotalStockLT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalStock), v))
	})
}

// TotalStockLTE applies the LTE predicate on the "totalStock" field.
func TotalStockLTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalStock), v))
	})
}

// ConsumeStockEQ applies the EQ predicate on the "consumeStock" field.
func ConsumeStockEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConsumeStock), v))
	})
}

// ConsumeStockNEQ applies the NEQ predicate on the "consumeStock" field.
func ConsumeStockNEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConsumeStock), v))
	})
}

// ConsumeStockIn applies the In predicate on the "consumeStock" field.
func ConsumeStockIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConsumeStock), v...))
	})
}

// ConsumeStockNotIn applies the NotIn predicate on the "consumeStock" field.
func ConsumeStockNotIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConsumeStock), v...))
	})
}

// ConsumeStockGT applies the GT predicate on the "consumeStock" field.
func ConsumeStockGT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConsumeStock), v))
	})
}

// ConsumeStockGTE applies the GTE predicate on the "consumeStock" field.
func ConsumeStockGTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConsumeStock), v))
	})
}

// ConsumeStockLT applies the LT predicate on the "consumeStock" field.
func ConsumeStockLT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConsumeStock), v))
	})
}

// ConsumeStockLTE applies the LTE predicate on the "consumeStock" field.
func ConsumeStockLTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConsumeStock), v))
	})
}

// LeftStockEQ applies the EQ predicate on the "leftStock" field.
func LeftStockEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeftStock), v))
	})
}

// LeftStockNEQ applies the NEQ predicate on the "leftStock" field.
func LeftStockNEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLeftStock), v))
	})
}

// LeftStockIn applies the In predicate on the "leftStock" field.
func LeftStockIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLeftStock), v...))
	})
}

// LeftStockNotIn applies the NotIn predicate on the "leftStock" field.
func LeftStockNotIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLeftStock), v...))
	})
}

// LeftStockGT applies the GT predicate on the "leftStock" field.
func LeftStockGT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLeftStock), v))
	})
}

// LeftStockGTE applies the GTE predicate on the "leftStock" field.
func LeftStockGTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLeftStock), v))
	})
}

// LeftStockLT applies the LT predicate on the "leftStock" field.
func LeftStockLT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLeftStock), v))
	})
}

// LeftStockLTE applies the LTE predicate on the "leftStock" field.
func LeftStockLTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLeftStock), v))
	})
}

// AddtimeEQ applies the EQ predicate on the "addtime" field.
func AddtimeEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddtime), v))
	})
}

// AddtimeNEQ applies the NEQ predicate on the "addtime" field.
func AddtimeNEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddtime), v))
	})
}

// AddtimeIn applies the In predicate on the "addtime" field.
func AddtimeIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddtime), v...))
	})
}

// AddtimeNotIn applies the NotIn predicate on the "addtime" field.
func AddtimeNotIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddtime), v...))
	})
}

// AddtimeGT applies the GT predicate on the "addtime" field.
func AddtimeGT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddtime), v))
	})
}

// AddtimeGTE applies the GTE predicate on the "addtime" field.
func AddtimeGTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddtime), v))
	})
}

// AddtimeLT applies the LT predicate on the "addtime" field.
func AddtimeLT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddtime), v))
	})
}

// AddtimeLTE applies the LTE predicate on the "addtime" field.
func AddtimeLTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddtime), v))
	})
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMtime), v))
	})
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMtime), v...))
	})
}

// MtimeNotIn applies the NotIn predicate on the "mtime" field.
func MtimeNotIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMtime), v...))
	})
}

// MtimeGT applies the GT predicate on the "mtime" field.
func MtimeGT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMtime), v))
	})
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMtime), v))
	})
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMtime), v))
	})
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMtime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		p(s.Not())
	})
}