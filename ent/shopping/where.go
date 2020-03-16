// Code generated by entc, DO NOT EDIT.

package shopping

import (
	"time"

	"github.com/Frosin/shoplist-api-client-go/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Shopping {
	return predicate.Shopping(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
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
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
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
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	},
	)
}

// Sum applies equality check predicate on the "sum" field. It's identical to SumEQ.
func Sum(v int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSum), v))
	},
	)
}

// Complete applies equality check predicate on the "complete" field. It's identical to CompleteEQ.
func Complete(v bool) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComplete), v))
	},
	)
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	},
	)
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	},
	)
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Shopping {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Shopping(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	},
	)
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Shopping {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Shopping(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	},
	)
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	},
	)
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	},
	)
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	},
	)
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	},
	)
}

// SumEQ applies the EQ predicate on the "sum" field.
func SumEQ(v int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSum), v))
	},
	)
}

// SumNEQ applies the NEQ predicate on the "sum" field.
func SumNEQ(v int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSum), v))
	},
	)
}

// SumIn applies the In predicate on the "sum" field.
func SumIn(vs ...int) predicate.Shopping {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Shopping(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSum), v...))
	},
	)
}

// SumNotIn applies the NotIn predicate on the "sum" field.
func SumNotIn(vs ...int) predicate.Shopping {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Shopping(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSum), v...))
	},
	)
}

// SumGT applies the GT predicate on the "sum" field.
func SumGT(v int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSum), v))
	},
	)
}

// SumGTE applies the GTE predicate on the "sum" field.
func SumGTE(v int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSum), v))
	},
	)
}

// SumLT applies the LT predicate on the "sum" field.
func SumLT(v int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSum), v))
	},
	)
}

// SumLTE applies the LTE predicate on the "sum" field.
func SumLTE(v int) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSum), v))
	},
	)
}

// CompleteEQ applies the EQ predicate on the "complete" field.
func CompleteEQ(v bool) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComplete), v))
	},
	)
}

// CompleteNEQ applies the NEQ predicate on the "complete" field.
func CompleteNEQ(v bool) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComplete), v))
	},
	)
}

// HasItem applies the HasEdge predicate on the "item" edge.
func HasItem() predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasItemWith applies the HasEdge predicate on the "item" edge with a given conditions (other predicates).
func HasItemWith(preds ...predicate.Item) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// HasShop applies the HasEdge predicate on the "shop" edge.
func HasShop() predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasShopWith applies the HasEdge predicate on the "shop" edge with a given conditions (other predicates).
func HasShopWith(preds ...predicate.Shop) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Shopping {
	return predicate.Shopping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Shopping) predicate.Shopping {
	return predicate.Shopping(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Shopping) predicate.Shopping {
	return predicate.Shopping(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Shopping) predicate.Shopping {
	return predicate.Shopping(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}