// Code generated by entc, DO NOT EDIT.

package deal

import (
	"kay/app/deal/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SkuId applies equality check predicate on the "skuId" field. It's identical to SkuIdEQ.
func SkuId(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuId), v))
	})
}

// ProductId applies equality check predicate on the "productId" field. It's identical to ProductIdEQ.
func ProductId(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductId), v))
	})
}

// ProductName applies equality check predicate on the "productName" field. It's identical to ProductNameEQ.
func ProductName(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductName), v))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// CustomName applies equality check predicate on the "customName" field. It's identical to CustomNameEQ.
func CustomName(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomName), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Addtime applies equality check predicate on the "addtime" field. It's identical to AddtimeEQ.
func Addtime(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddtime), v))
	})
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// SkuIdEQ applies the EQ predicate on the "skuId" field.
func SkuIdEQ(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuId), v))
	})
}

// SkuIdNEQ applies the NEQ predicate on the "skuId" field.
func SkuIdNEQ(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkuId), v))
	})
}

// SkuIdIn applies the In predicate on the "skuId" field.
func SkuIdIn(vs ...int64) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkuId), v...))
	})
}

// SkuIdNotIn applies the NotIn predicate on the "skuId" field.
func SkuIdNotIn(vs ...int64) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkuId), v...))
	})
}

// SkuIdGT applies the GT predicate on the "skuId" field.
func SkuIdGT(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkuId), v))
	})
}

// SkuIdGTE applies the GTE predicate on the "skuId" field.
func SkuIdGTE(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkuId), v))
	})
}

// SkuIdLT applies the LT predicate on the "skuId" field.
func SkuIdLT(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkuId), v))
	})
}

// SkuIdLTE applies the LTE predicate on the "skuId" field.
func SkuIdLTE(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkuId), v))
	})
}

// ProductIdEQ applies the EQ predicate on the "productId" field.
func ProductIdEQ(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductId), v))
	})
}

// ProductIdNEQ applies the NEQ predicate on the "productId" field.
func ProductIdNEQ(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductId), v))
	})
}

// ProductIdIn applies the In predicate on the "productId" field.
func ProductIdIn(vs ...int64) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductId), v...))
	})
}

// ProductIdNotIn applies the NotIn predicate on the "productId" field.
func ProductIdNotIn(vs ...int64) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductId), v...))
	})
}

// ProductIdGT applies the GT predicate on the "productId" field.
func ProductIdGT(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductId), v))
	})
}

// ProductIdGTE applies the GTE predicate on the "productId" field.
func ProductIdGTE(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductId), v))
	})
}

// ProductIdLT applies the LT predicate on the "productId" field.
func ProductIdLT(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductId), v))
	})
}

// ProductIdLTE applies the LTE predicate on the "productId" field.
func ProductIdLTE(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductId), v))
	})
}

// ProductNameEQ applies the EQ predicate on the "productName" field.
func ProductNameEQ(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductName), v))
	})
}

// ProductNameNEQ applies the NEQ predicate on the "productName" field.
func ProductNameNEQ(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductName), v))
	})
}

// ProductNameIn applies the In predicate on the "productName" field.
func ProductNameIn(vs ...string) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductName), v...))
	})
}

// ProductNameNotIn applies the NotIn predicate on the "productName" field.
func ProductNameNotIn(vs ...string) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductName), v...))
	})
}

// ProductNameGT applies the GT predicate on the "productName" field.
func ProductNameGT(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductName), v))
	})
}

// ProductNameGTE applies the GTE predicate on the "productName" field.
func ProductNameGTE(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductName), v))
	})
}

// ProductNameLT applies the LT predicate on the "productName" field.
func ProductNameLT(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductName), v))
	})
}

// ProductNameLTE applies the LTE predicate on the "productName" field.
func ProductNameLTE(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductName), v))
	})
}

// ProductNameContains applies the Contains predicate on the "productName" field.
func ProductNameContains(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProductName), v))
	})
}

// ProductNameHasPrefix applies the HasPrefix predicate on the "productName" field.
func ProductNameHasPrefix(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProductName), v))
	})
}

// ProductNameHasSuffix applies the HasSuffix predicate on the "productName" field.
func ProductNameHasSuffix(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProductName), v))
	})
}

// ProductNameEqualFold applies the EqualFold predicate on the "productName" field.
func ProductNameEqualFold(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProductName), v))
	})
}

// ProductNameContainsFold applies the ContainsFold predicate on the "productName" field.
func ProductNameContainsFold(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProductName), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...int64) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUID), v...))
	})
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...int64) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUID), v...))
	})
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v int64) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// CustomNameEQ applies the EQ predicate on the "customName" field.
func CustomNameEQ(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomName), v))
	})
}

// CustomNameNEQ applies the NEQ predicate on the "customName" field.
func CustomNameNEQ(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustomName), v))
	})
}

// CustomNameIn applies the In predicate on the "customName" field.
func CustomNameIn(vs ...string) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCustomName), v...))
	})
}

// CustomNameNotIn applies the NotIn predicate on the "customName" field.
func CustomNameNotIn(vs ...string) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCustomName), v...))
	})
}

// CustomNameGT applies the GT predicate on the "customName" field.
func CustomNameGT(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCustomName), v))
	})
}

// CustomNameGTE applies the GTE predicate on the "customName" field.
func CustomNameGTE(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCustomName), v))
	})
}

// CustomNameLT applies the LT predicate on the "customName" field.
func CustomNameLT(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCustomName), v))
	})
}

// CustomNameLTE applies the LTE predicate on the "customName" field.
func CustomNameLTE(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCustomName), v))
	})
}

// CustomNameContains applies the Contains predicate on the "customName" field.
func CustomNameContains(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCustomName), v))
	})
}

// CustomNameHasPrefix applies the HasPrefix predicate on the "customName" field.
func CustomNameHasPrefix(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCustomName), v))
	})
}

// CustomNameHasSuffix applies the HasSuffix predicate on the "customName" field.
func CustomNameHasSuffix(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCustomName), v))
	})
}

// CustomNameEqualFold applies the EqualFold predicate on the "customName" field.
func CustomNameEqualFold(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCustomName), v))
	})
}

// CustomNameContainsFold applies the ContainsFold predicate on the "customName" field.
func CustomNameContainsFold(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCustomName), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// AddtimeEQ applies the EQ predicate on the "addtime" field.
func AddtimeEQ(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddtime), v))
	})
}

// AddtimeNEQ applies the NEQ predicate on the "addtime" field.
func AddtimeNEQ(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddtime), v))
	})
}

// AddtimeIn applies the In predicate on the "addtime" field.
func AddtimeIn(vs ...time.Time) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
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
func AddtimeNotIn(vs ...time.Time) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
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
func AddtimeGT(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddtime), v))
	})
}

// AddtimeGTE applies the GTE predicate on the "addtime" field.
func AddtimeGTE(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddtime), v))
	})
}

// AddtimeLT applies the LT predicate on the "addtime" field.
func AddtimeLT(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddtime), v))
	})
}

// AddtimeLTE applies the LTE predicate on the "addtime" field.
func AddtimeLTE(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddtime), v))
	})
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMtime), v))
	})
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
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
func MtimeNotIn(vs ...time.Time) predicate.Deal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deal(func(s *sql.Selector) {
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
func MtimeGT(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMtime), v))
	})
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMtime), v))
	})
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMtime), v))
	})
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMtime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Deal) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Deal) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
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
func Not(p predicate.Deal) predicate.Deal {
	return predicate.Deal(func(s *sql.Selector) {
		p(s.Not())
	})
}
