package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreatedBy struct{ mixin.Schema }

func (CreatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("created_by").
			Comment("创建者ID").
			Optional().
			Nillable(),
	}
}

type UpdatedBy struct{ mixin.Schema }

func (UpdatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("updated_by").
			Comment("更新者ID").
			Optional().
			Nillable(),
	}
}

type DeletedBy struct{ mixin.Schema }

func (DeletedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("deleted_by").
			Comment("删除者ID").
			Optional().
			Nillable(),
	}
}
