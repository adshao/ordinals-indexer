package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Inscription is a Inscription model.
type Inscription struct {
	ID            int       `json:"id"`
	InscriptionID int64     `json:"inscription_id"`
	UID           string    `json:"uid"`
	Address       string    `json:"address"`
	OutputValue   uint64    `json:"output_value"`
	ContentLength uint64    `json:"content_length"`
	ContentType   string    `json:"content_type"`
	Timestamp     time.Time `json:"timestamp"`
	GenesisHeight uint64    `json:"genesis_height"`
	GenesisFee    uint64    `json:"genesis_fee"`
	GenesisTx     string    `json:"genesis_tx"`
	Location      string    `json:"location"`
	Output        string    `json:"output"`
	Offset        uint64    `json:"offset"`
}

type InscriptionListOption struct {
	Limit  int
	Offset int
	Order  string
}

// InscriptionRepo is a Greater repo.
type InscriptionRepo interface {
	Create(context.Context, *Inscription) (*Inscription, error)
	Update(context.Context, *Inscription) (*Inscription, error)
	FindByInscriptionID(context.Context, int64) (*Inscription, error)
	List(context.Context, ...InscriptionListOption) ([]*Inscription, error)
	Delete(context.Context, int) error
	Count(context.Context, ...InscriptionListOption) (int, error)
}

// InscriptionUsecase is a Inscription usecase.
type InscriptionUsecase struct {
	repo InscriptionRepo
	log  *log.Helper
}

// NewInscriptionUsecase new a Inscription usecase.
func NewInscriptionUsecase(repo InscriptionRepo, logger log.Logger) *InscriptionUsecase {
	return &InscriptionUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateInscription creates a Inscription, and returns the new Inscription.
func (uc *InscriptionUsecase) CreateInscription(ctx context.Context, g *Inscription) (*Inscription, error) {
	uc.log.WithContext(ctx).Debugf("CreateInscription for inscription %d", g.InscriptionID)
	return uc.repo.Create(ctx, g)
}

// UpdateInscription updates a Inscription, and returns the new Inscription.
func (uc *InscriptionUsecase) UpdateInscription(ctx context.Context, g *Inscription) (*Inscription, error) {
	uc.log.WithContext(ctx).Debugf("UpdateInscription for inscription %d", g.InscriptionID)
	return uc.repo.Update(ctx, g)
}

// FindByInscriptionID finds the Inscription by InscriptionID.
func (uc *InscriptionUsecase) FindByInscriptionID(ctx context.Context, inscriptionID int64) (*Inscription, error) {
	uc.log.WithContext(ctx).Debugf("FindByInscriptionID for %d", inscriptionID)
	return uc.repo.FindByInscriptionID(ctx, inscriptionID)
}

// ListInscriptions lists Inscriptions.
func (uc *InscriptionUsecase) ListInscriptions(ctx context.Context, opt *InscriptionListOption) ([]*Inscription, error) {
	uc.log.WithContext(ctx).Debugf("ListInscriptions for %v", opt)
	return uc.repo.List(ctx, *opt)
}

// DeleteInscription deletes a Inscription.
func (uc *InscriptionUsecase) DeleteInscription(ctx context.Context, id int) error {
	uc.log.WithContext(ctx).Debugf("DeleteInscription for %d", id)
	return uc.repo.Delete(ctx, id)
}

// CountInscriptions counts Inscriptions.
func (uc *InscriptionUsecase) CountInscriptions(ctx context.Context, opt *InscriptionListOption) (int, error) {
	uc.log.WithContext(ctx).Debugf("CountInscriptions for %v", opt)
	return uc.repo.Count(ctx, *opt)
}
