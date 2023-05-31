package data

import (
	"context"
	"strings"

	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/data/ent"
	"github.com/adshao/ordinals-indexer/internal/data/ent/token"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	defaultListLimit = 100
)

type tokenRepo struct {
	data                 *Data
	log                  *log.Helper
	orderFieldsWhiteList []string
}

// NewTokenRepo .
func NewTokenRepo(data *Data, logger log.Logger) biz.TokenRepo {
	return &tokenRepo{
		data:                 data,
		log:                  log.NewHelper(logger),
		orderFieldsWhiteList: []string{"id", "created_at", "p", "tick", "token_id", "block_height", "block_time", "inscription_id"},
	}
}

func (r *tokenRepo) Create(ctx context.Context, g *biz.Token) (*biz.Token, error) {
	res, err := r.data.db.Token.Create().
		SetP(g.P).
		SetTick(g.Tick).
		SetTokenID(g.TokenID).
		SetTxHash(g.TxHash).
		SetBlockHeight(g.BlockHeight).
		SetInscriptionID(g.InscriptionID).
		SetInscriptionUID(g.InscriptionUID).
		SetCollection(&ent.Collection{ID: g.CollectionID}).
		SetBlockTime(g.BlockTime).
		SetAddress(g.Address).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.fromDbToken(res), nil
}

func (r *tokenRepo) fromDbToken(t *ent.Token) *biz.Token {
	token := &biz.Token{
		ID:             t.ID,
		Tick:           t.Tick,
		P:              t.P,
		TokenID:        t.TokenID,
		TxHash:         t.TxHash,
		BlockHeight:    t.BlockHeight,
		BlockTime:      t.BlockTime,
		Address:        t.Address,
		InscriptionID:  t.InscriptionID,
		InscriptionUID: t.InscriptionUID,
	}
	if t.Edges.Collection != nil {
		token.CollectionID = t.Edges.Collection.ID
	}
	return token
}

func (r *tokenRepo) Update(ctx context.Context, g *biz.Token) (*biz.Token, error) {
	u := r.data.db.Token.UpdateOneID(g.ID).
		SetP(g.P).
		SetTick(g.Tick).
		SetTokenID(g.TokenID).
		SetTxHash(g.TxHash).
		SetBlockHeight(g.BlockHeight).
		SetInscriptionID(g.InscriptionID).
		SetInscriptionUID(g.InscriptionUID).
		SetBlockTime(g.BlockTime).
		SetAddress(g.Address)
	if g.CollectionID != 0 {
		u.SetCollection(&ent.Collection{ID: g.CollectionID})
	}
	res, err := u.Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.fromDbToken(res), nil
}

func (r *tokenRepo) FindByTickTokenID(ctx context.Context, p, tick string, tokenID uint64) (*biz.Token, error) {
	res, err := r.data.db.Token.Query().Where(token.P(p), token.Tick(tick), token.TokenID(tokenID)).WithCollection().Only(ctx)
	if err == nil {
		return r.fromDbToken(res), nil
	}
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return nil, err
}

func (r *tokenRepo) FindByInscriptionID(ctx context.Context, id int64) ([]*biz.Token, error) {
	res, err := r.data.db.Token.Query().Where(token.InscriptionID(id)).WithCollection().All(ctx)
	if err != nil {
		return nil, err
	}
	var ret []*biz.Token
	for _, token := range res {
		ret = append(ret, r.fromDbToken(token))
	}
	return ret, nil
}

func (r *tokenRepo) List(ctx context.Context, opts ...biz.TokenListOption) ([]*biz.Token, error) {
	q := r.data.db.Token.Query()
	var opt biz.TokenListOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Limit > 0 && opt.Limit <= defaultListLimit {
		q = q.Limit(opt.Limit)
	} else {
		q = q.Limit(defaultListLimit)
	}
	if opt.Offset != 0 {
		q = q.Offset(opt.Offset)
	}
	if opt.P != "" {
		q = q.Where(token.P(opt.P))
	}
	if opt.Tick != "" {
		q = q.Where(token.Tick(opt.Tick))
	}
	// order format: field1,-field2
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

	res, err := q.WithCollection().All(ctx)
	if err != nil {
		return nil, err
	}
	var ret []*biz.Token
	for _, token := range res {
		ret = append(ret, r.fromDbToken(token))
	}
	return ret, nil
}

func (r *tokenRepo) inOrderFieldsWhiteList(field string) bool {
	for _, f := range r.orderFieldsWhiteList {
		if f == field {
			return true
		}
	}
	return false
}

func (r *tokenRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.Token.DeleteOneID(id).Exec(ctx)
}

func (r *tokenRepo) Count(ctx context.Context, opts ...biz.TokenListOption) (int, error) {
	q := r.data.db.Token.Query()
	var opt biz.TokenListOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.P != "" {
		q = q.Where(token.P(opt.P))
	}
	if opt.Tick != "" {
		q = q.Where(token.Tick(opt.Tick))
	}
	return q.Count(ctx)
}
