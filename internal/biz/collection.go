package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Collection is a Collection model.
type Collection struct {
	ID             int                      `json:"id"`
	Tick           string                   `json:"tick"`
	P              string                   `json:"p"`
	Max            uint64                   `json:"max"`
	Supply         uint64                   `json:"supply,omitempty"`
	BaseURI        string                   `json:"base_uri,omitempty"`
	Name           string                   `json:"name,omitempty"`
	Description    string                   `json:"description,omitempty"`
	Image          string                   `json:"image,omitempty"`
	Attributes     []map[string]interface{} `json:"attributes,omitempty"`
	TxHash         string                   `json:"tx_hash,omitempty"`
	BlockHeight    uint64                   `json:"block_height"`
	BlockTime      time.Time                `json:"block_time"`
	Address        string                   `json:"address,omitempty"`
	InscriptionID  int64                    `json:"inscription_id"`
	InscriptionUID string                   `json:"inscription_uid"`
}

type CollectionListOption struct {
	Limit  int
	Offset int
	P      string
	Tick   string
	Order  string
}

// CollectionRepo is a Greater repo.
type CollectionRepo interface {
	Create(context.Context, *Collection) (*Collection, error)
	Update(context.Context, *Collection) (*Collection, error)
	FindByID(context.Context, int) (*Collection, error)
	FindByTick(context.Context, string, string) (*Collection, error)
	FindByInscriptionID(context.Context, int64) ([]*Collection, error)
	List(context.Context, ...CollectionListOption) ([]*Collection, error)
	Delete(context.Context, int) error
	Count(context.Context, ...CollectionListOption) (int, error)
}

// CollectionUsecase is a Collection usecase.
type CollectionUsecase struct {
	repo CollectionRepo
	log  *log.Helper
}

// NewCollectionUsecase new a Collection usecase.
func NewCollectionUsecase(repo CollectionRepo, logger log.Logger) *CollectionUsecase {
	return &CollectionUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCollection creates a Collection, and returns the new Collection.
func (uc *CollectionUsecase) CreateCollection(ctx context.Context, g *Collection) (*Collection, error) {
	uc.log.WithContext(ctx).Debugf("CreateCollection for inscription %d", g.InscriptionID)
	return uc.repo.Create(ctx, g)
}

// UpdateCollection updates a Collection, and returns the new Collection.
func (uc *CollectionUsecase) UpdateCollection(ctx context.Context, g *Collection) (*Collection, error) {
	uc.log.WithContext(ctx).Debugf("UpdateCollection for inscription %d", g.InscriptionID)
	return uc.repo.Update(ctx, g)
}

// GetCollection gets a Collection by the ID.
func (uc *CollectionUsecase) GetCollection(ctx context.Context, id int) (*Collection, error) {
	collection, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (uc *CollectionUsecase) GetCollectionByInscriptionID(ctx context.Context, inscriptionID int64) ([]*Collection, error) {
	collections, err := uc.repo.FindByInscriptionID(ctx, inscriptionID)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (uc *CollectionUsecase) GetCollectionByTick(ctx context.Context, p string, tick string) (*Collection, error) {
	collection, err := uc.repo.FindByTick(ctx, p, tick)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

// ListCollections list all Collections.
func (uc *CollectionUsecase) ListCollections(ctx context.Context, opt *CollectionListOption) ([]*Collection, error) {
	return uc.repo.List(ctx, *opt)
}

// DeleteCollection deletes a Collection.
func (uc *CollectionUsecase) DeleteCollection(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

// CountCollection counts the number of the Collection.
func (uc *CollectionUsecase) CountCollection(ctx context.Context, opt *CollectionListOption) (int, error) {
	return uc.repo.Count(ctx, *opt)
}
