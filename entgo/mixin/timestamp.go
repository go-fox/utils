package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CreatedTimestamp 创建时间
type CreatedTimestamp struct{ mixin.Schema }

func (CreatedTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间，毫秒
		field.Int64("created_at").
			Comment("创建时间").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().Unix()),
	}
}

// UpdatedTimestamp 更新时间
type UpdatedTimestamp struct{ mixin.Schema }

func (UpdatedTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间，毫秒
		field.Int64("updated_at").
			Comment("修改时间").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().Unix()),
	}
}

// DeletedTimestamp 删除时间
type DeletedTimestamp struct{ mixin.Schema }

func (DeletedTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// 删除时间，毫秒
		field.Int64("deleted_at").
			Comment("删除时间").
			Optional().
			Nillable(),
	}
}
