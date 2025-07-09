package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Status 状态字段
type Status struct {
	mixin.Schema
}

// Fields of the Status.
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

// IntegerStatus 整数状态字段
type IntegerStatus struct {
	mixin.Schema
}

// Fields of the IntegerStatus.
func (IntegerStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("status").
			Max(3).
			Default(1).
			Comment("状态").
			Optional().
			Nillable(),
	}
}
