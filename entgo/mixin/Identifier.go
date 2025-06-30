package mixin

import (
	"regexp"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type AutoIncrementId struct{ mixin.Schema }

func (AutoIncrementId) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Comment("id").
			StructTag(`json:"id,omitempty"`).
			SchemaType(map[string]string{
				dialect.MySQL:    "int",
				dialect.Postgres: "serial",
			}).
			Annotations(
				entproto.Field(1),
			).
			Positive().
			Immutable().
			Unique(),
	}
}

// Indexes of the AutoIncrementId.
func (AutoIncrementId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

// UUID uuid
type UUID struct{ mixin.Schema }

// Fields of the UUID.
func (UUID) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("id").
			Default(uuid.New).
			Unique().
			Immutable(),
	}
}

// Indexes of the UuidId.
func (UUID) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

// StringId string id
type StringId struct {
	mixin.Schema
}

func (StringId) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Comment("id").
			MaxLen(50).
			NotEmpty().
			Unique().
			Immutable().
			Match(regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")),
	}
}

// Indexes of the StringId.
func (StringId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
