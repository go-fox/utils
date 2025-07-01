package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Remark 备注
type Remark struct {
	mixin.Schema
}

// Fields of the CreateBy.
func (Remark) Fields() []ent.Field {
	return []ent.Field{
		field.Text("remark").
			Comment("备注").
			Optional().
			Nillable(),
	}
}
