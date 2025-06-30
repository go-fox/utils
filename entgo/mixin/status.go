package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	_mixin "entgo.io/ent/schema/mixin"
)

// Status 状态字段
type Status struct {
	_mixin.Schema
}

// Fields of the CreateBy.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("enabled", "disabled").
			Default("enabled").
			Comment("状态").
			Optional().
			Nillable(),
	}
}
