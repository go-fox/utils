package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Sort 排序
type Sort struct {
	mixin.Schema
}

// Fields of the CreateBy.
func (Sort) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("sort").
			Comment("排序").
			Optional().
			Nillable().
			Default(0),
	}
}
