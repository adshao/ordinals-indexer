package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Inscription holds the schema definition for the Inscription entity.
type Inscription struct {
	ent.Schema
}

func (Inscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Inscription.
func (Inscription) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("inscription_id").Unique(),
		field.String("uid").Unique(),
		field.String("address"),
		field.Uint64("output_value"),
		field.Uint64("content_length"),
		field.String("content_type"),
		field.Time("timestamp"),
		field.Uint64("genesis_height"),
		field.Uint64("genesis_fee"),
		field.String("genesis_tx"),
		field.String("location"),
		field.String("output"),
		field.Uint64("offset"),
	}
}

// Edges of the Inscription.
func (Inscription) Edges() []ent.Edge {
	return nil
}
