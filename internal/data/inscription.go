package data

import (
	"context"
	"strings"

	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/data/ent"
	"github.com/adshao/ordinals-indexer/internal/data/ent/inscription"

	"github.com/go-kratos/kratos/v2/log"
)

type inscriptionRepo struct {
	data                 *Data
	log                  *log.Helper
	orderFieldsWhiteList []string
}

// NewInscriptionRepo .
func NewInscriptionRepo(data *Data, logger log.Logger) biz.InscriptionRepo {
	return &inscriptionRepo{
		data:                 data,
		log:                  log.NewHelper(logger),
		orderFieldsWhiteList: []string{"id", "created_at", "inscription_id", "timestamp"},
	}
}

func (r *inscriptionRepo) Create(ctx context.Context, g *biz.Inscription) (*biz.Inscription, error) {
	res, err := r.data.db.Inscription.Create().
		SetInscriptionID(g.InscriptionID).
		SetUID(g.UID).
		SetAddress(g.Address).
		SetOutput(g.Output).
		SetContentLength(g.ContentLength).
		SetContentType(g.ContentType).
		SetTimestamp(g.Timestamp).
		SetGenesisHeight(g.GenesisHeight).
		SetGenesisFee(g.GenesisFee).
		SetGenesisTx(g.GenesisTx).
		SetLocation(g.Location).
		SetOutput(g.Output).
		SetOffset(g.Offset).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.fromDbInscription(res), nil
}

func (r *inscriptionRepo) fromDbInscription(t *ent.Inscription) *biz.Inscription {
	token := &biz.Inscription{
		InscriptionID: t.InscriptionID,
		UID:           t.UID,
		Address:       t.Address,
		OutputValue:   t.OutputValue,
		ContentLength: t.ContentLength,
		ContentType:   t.ContentType,
		Timestamp:     t.Timestamp,
		GenesisHeight: t.GenesisHeight,
		GenesisFee:    t.GenesisFee,
		GenesisTx:     t.GenesisTx,
		Location:      t.Location,
		Output:        t.Output,
		Offset:        t.Offset,
	}
	return token
}

func (r *inscriptionRepo) Update(ctx context.Context, g *biz.Inscription) (*biz.Inscription, error) {
	u := r.data.db.Inscription.UpdateOneID(g.ID).
		SetInscriptionID(g.InscriptionID).
		SetUID(g.UID).
		SetAddress(g.Address).
		SetOutput(g.Output).
		SetContentLength(g.ContentLength).
		SetContentType(g.ContentType).
		SetTimestamp(g.Timestamp).
		SetGenesisHeight(g.GenesisHeight).
		SetGenesisFee(g.GenesisFee).
		SetGenesisTx(g.GenesisTx).
		SetLocation(g.Location).
		SetOutput(g.Output).
		SetOffset(g.Offset)
	res, err := u.Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.fromDbInscription(res), nil
}

func (r *inscriptionRepo) FindByInscriptionID(ctx context.Context, inscriptionID int64) (*biz.Inscription, error) {
	res, err := r.data.db.Inscription.Query().Where(inscription.InscriptionID(inscriptionID)).Only(ctx)
	if err == nil {
		return r.fromDbInscription(res), nil
	}
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return nil, err
}

func (r *inscriptionRepo) List(ctx context.Context, opts ...biz.InscriptionListOption) ([]*biz.Inscription, error) {
	q := r.data.db.Inscription.Query()
	var opt biz.InscriptionListOption
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

	res, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	var ret []*biz.Inscription
	for _, ins := range res {
		ret = append(ret, r.fromDbInscription(ins))
	}
	return ret, nil
}

func (r *inscriptionRepo) inOrderFieldsWhiteList(field string) bool {
	for _, f := range r.orderFieldsWhiteList {
		if f == field {
			return true
		}
	}
	return false
}

func (r *inscriptionRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.Inscription.DeleteOneID(id).Exec(ctx)
}

func (r *inscriptionRepo) Count(ctx context.Context, opts ...biz.InscriptionListOption) (int, error) {
	q := r.data.db.Inscription.Query()
	return q.Count(ctx)
}
