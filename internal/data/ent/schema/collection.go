package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

func (Collection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.String("tick"),
		field.String("p").Default("brc-721"),
		field.Uint64("max"),
		field.Uint64("supply"),
		field.String("base_uri"),
		field.String("name"),
		field.String("description"),
		field.String("image"),
		field.JSON("attributes", []map[string]interface{}{}),
		field.String("tx_hash"),
		field.Uint64("block_height"),
		field.Time("block_time"),
		field.String("address"),
		field.Int64("inscription_id").Unique(),
		field.String("inscription_uid").Unique(),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", Token.Type),
	}
}

func (Collection) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("p", "tick").Unique(),
		index.Fields("tx_hash"),
		index.Fields("block_height"),
		index.Fields("inscription_id"),
		index.Fields("address"),
	}
}
