package data

import (
	"context"
	"strings"

	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/data/ent"
	"github.com/adshao/ordinals-indexer/internal/data/ent/collection"

	"github.com/go-kratos/kratos/v2/log"
)

type collectionRepo struct {
	data                 *Data
	log                  *log.Helper
	orderFieldsWhiteList []string
}

// NewCollectionRepo .
func NewCollectionRepo(data *Data, logger log.Logger) biz.CollectionRepo {
	return &collectionRepo{
		data:                 data,
		log:                  log.NewHelper(logger),
		orderFieldsWhiteList: []string{"id", "created_at", "p", "tick", "block_height", "block_time", "inscription_id"},
	}
}

func (r *collectionRepo) Create(ctx context.Context, g *biz.Collection) (*biz.Collection, error) {
	res, err := r.data.db.Collection.Create().
		SetP(g.P).
		SetTick(g.Tick).
		SetMax(g.Max).
		SetSupply(g.Supply).
		SetBaseURI(g.BaseURI).
		SetName(g.Name).
		SetDescription(g.Description).
		SetImage(g.Image).
		SetAttributes(g.Attributes).
		SetTxHash(g.TxHash).
		SetBlockHeight(g.BlockHeight).
		SetBlockTime(g.BlockTime).
		SetAddress(g.Address).
		SetInscriptionID(g.InscriptionID).
		SetInscriptionUID(g.InscriptionUID).
		SetSig(g.Sig).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.fromDbCollection(res), err
}

func (r *collectionRepo) fromDbCollection(t *ent.Collection) *biz.Collection {
	collection := &biz.Collection{
		ID:             t.ID,
		P:              t.P,
		Tick:           t.Tick,
		Max:            t.Max,
		Supply:         t.Supply,
		BaseURI:        t.BaseURI,
		Name:           t.Name,
		Description:    t.Description,
		Image:          t.Image,
		Attributes:     t.Attributes,
		TxHash:         t.TxHash,
		BlockHeight:    t.BlockHeight,
		BlockTime:      t.BlockTime,
		Address:        t.Address,
		InscriptionID:  t.InscriptionID,
		InscriptionUID: t.InscriptionUID,
		Sig:            t.Sig,
	}
	return collection
}

func (r *collectionRepo) Update(ctx context.Context, g *biz.Collection) (*biz.Collection, error) {
	res, err := r.data.db.Collection.UpdateOneID(g.ID).
		SetP(g.P).
		SetTick(g.Tick).
		SetMax(g.Max).
		SetSupply(g.Supply).
		SetBaseURI(g.BaseURI).
		SetName(g.Name).
		SetDescription(g.Description).
		SetImage(g.Image).
		SetAttributes(g.Attributes).
		SetTxHash(g.TxHash).
		SetBlockHeight(g.BlockHeight).
		SetBlockTime(g.BlockTime).
		SetAddress(g.Address).
		SetInscriptionID(g.InscriptionID).
		SetInscriptionUID(g.InscriptionUID).
		SetSig(g.Sig).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.fromDbCollection(res), err
}

func (r *collectionRepo) FindByID(ctx context.Context, id int) (*biz.Collection, error) {
	res, err := r.data.db.Collection.Query().Where(collection.IDEQ(id)).Only(ctx)
	if err == nil {
		return r.fromDbCollection(res), nil
	}
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return nil, err
}

func (r *collectionRepo) FindByTick(ctx context.Context, p, tick string) (*biz.Collection, error) {
	res, err := r.data.db.Collection.Query().Where(collection.PEQ(p), collection.TickEQ(tick)).Only(ctx)
	if err == nil {
		return r.fromDbCollection(res), nil
	}
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return nil, err
}

func (r *collectionRepo) FindByInscriptionID(ctx context.Context, id int64) ([]*biz.Collection, error) {
	items := make([]*biz.Collection, 0)
	res, err := r.data.db.Collection.Query().Where(collection.InscriptionIDEQ(id)).All(ctx)
	if err == nil {
		for _, collection := range res {
			items = append(items, r.fromDbCollection(collection))
		}
		return items, nil
	}
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return nil, err
}

func (r *collectionRepo) List(ctx context.Context, opts ...biz.CollectionListOption) ([]*biz.Collection, error) {
	q := r.data.db.Collection.Query()
	var opt biz.CollectionListOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Limit != 0 && opt.Limit <= defaultListLimit {
		q = q.Limit(opt.Limit)
	} else {
		q = q.Limit(defaultListLimit)
	}
	if opt.Offset != 0 {
		q = q.Offset(opt.Offset)
	}
	if opt.P != "" {
		q = q.Where(collection.PEQ(opt.P))
	}
	if opt.Tick != "" {
		q = q.Where(collection.TickEQ(opt.Tick))
	}
	// order format: "id,created_at,-tick"
	if opt.Order != "" {
		orders := strings.Split(opt.Order, ",")
		for _, order := range orders {
			asc := true
			field := strings.ToLower(order)
			if strings.HasPrefix(order, "-") {
				field = strings.TrimPrefix(order, "-")
				asc = false
			}
			if !r.inOrderFieldsWhiteList(field) {
				continue
			}
			if asc {
				q = q.Order(ent.Asc(field))
			} else {
				q = q.Order(ent.Desc(field))
			}
		}
	}
	res, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*biz.Collection, 0)
	for _, collection := range res {
		items = append(items, r.fromDbCollection(collection))
	}
	return items, nil
}

func (r *collectionRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.Collection.DeleteOneID(id).Exec(ctx)
}

func (r *collectionRepo) Count(ctx context.Context, opts ...biz.CollectionListOption) (int, error) {
	q := r.data.db.Collection.Query()
	var opt biz.CollectionListOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.P != "" {
		q = q.Where(collection.PEQ(opt.P))
	}
	if opt.Tick != "" {
		q = q.Where(collection.TickEQ(opt.Tick))
	}
	return q.Count(ctx)
}

func (r *collectionRepo) inOrderFieldsWhiteList(field string) bool {
	for _, f := range r.orderFieldsWhiteList {
		if f == field {
			return true
		}
	}
	return false
}
