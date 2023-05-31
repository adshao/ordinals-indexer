package biz

import (
	"context"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCollectionUsecase, NewTokenUsecase, NewInscriptionUsecase)

type RedisRepo interface {
	GetLastInscriptionId(ctx context.Context) (int64, error)
	SetLastInscriptionId(ctx context.Context, id int64) error
}
