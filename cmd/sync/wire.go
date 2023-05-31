//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/data"
	"github.com/adshao/ordinals-indexer/internal/ord"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Ord, *conf.Data, log.Logger) (*ord.Syncer, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, newApp))
}
