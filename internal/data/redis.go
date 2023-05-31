package data

import (
	// "context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	lastInscriptionIdKey = "lastInscriptionId"
)

func collectionTickKey(tick string) string {
	return "collection:" + tick
}

type redisRepo struct {
	data *Data
	log  *log.Helper
}

// NewRedisRepo .
func NewRedisRepo(data *Data, logger log.Logger) *redisRepo {
	return &redisRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// func (r *redisRepo) GetLastInscriptionId(ctx context.Context) (uint64, error) {
// 	return r.data.rdb.GetUint64(ctx, lastInscriptionIdKey)
// }

// func (r *redisRepo) SetLastInscriptionId(ctx context.Context, id uint64) error {
// 	return r.data.rdb.SetUint64(ctx, lastInscriptionIdKey, id)
// }

// func (r *redisRepo) GetCollectionIDByTick(ctx context.Context, tick string) (uint64, error) {
// 	return r.data.rdb.GetUint64(ctx, collectionTickKey(tick))
// }

// func (r *redisRepo) SetCollectionIDByTick(ctx context.Context, tick string, id uint64) error {
// 	return r.data.rdb.SetUint64(ctx, collectionTickKey(tick), id)
// }
