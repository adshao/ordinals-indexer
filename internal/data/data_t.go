package data

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"

	"github.com/adshao/ordinals-indexer/internal/data/ent"
	"github.com/adshao/ordinals-indexer/internal/data/ent/enttest"
	"github.com/adshao/ordinals-indexer/internal/data/ent/migrate"
)

func NewTData(t *testing.T) (*Data, func()) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	return &Data{
			db: client,
		}, func() {
			client.Close()
		}
}
