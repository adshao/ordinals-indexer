package data

import (
	"context"
	"fmt"

	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/data/ent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	// init postgresql driver.
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTokenRepo, NewCollectionRepo, NewRedisRepo, NewInscriptionRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx).Debug(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		log.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
