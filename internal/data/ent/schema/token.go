package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/adshao/go-brc721/sig"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("tick"),
		field.String("p").Default("brc-721"),
		field.Uint64("token_id"),
		field.String("tx_hash"),
		field.Uint64("block_height"),
		field.Time("block_time"),
		field.String("address"),
		field.Int64("inscription_id").Unique(),
		field.String("inscription_uid").Unique(),
		field.JSON("sig", sig.MintSig{}).Optional(),
		field.String("sig_uid"),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("collection", Collection.Type).Ref("tokens").Unique(),
	}
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("p", "tick", "token_id").Unique(),
		index.Fields("address"),
		index.Fields("tx_hash"),
		index.Fields("inscription_id").Unique(),
		index.Fields("block_height"),
		index.Fields("sig_uid"),
	}
}
