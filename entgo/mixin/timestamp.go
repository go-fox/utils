package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CreateTimestamp 创建时间
type CreateTimestamp struct{ mixin.Schema }

func (CreateTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间，毫秒
		field.Int64("create_at").
			Comment("创建时间").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().Unix()),
	}
}

// UpdateTimestamp 更新时间
type UpdateTimestamp struct{ mixin.Schema }

func (UpdateTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间，毫秒
		field.Int64("update_at").
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
