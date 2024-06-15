package usecase

import (
	"VinylShop/internal/vinyls/domain"
	ctx "VinylShop/pkg/types/context"
	"VinylShop/pkg/types/responses"
	errs "VinylShop/pkg/web/errors"
	"context"
)

type VinylsUseCase interface {
	InsertVinyl(ctx context.Context, vinyl *domain.Vinyl) (int, *errs.Error)
	GetVinyl(ctx context.Context, id int) (*domain.Vinyl, *errs.Error)
	GetVinyls(ctx ctx.Context, filters *domain.Filters) (*responses.Pagination, *errs.Error)
	UpdateVinyl(ctx context.Context, id int, vinyl *domain.Vinyl) (*domain.Vinyl, *errs.Error)
	DeleteVinyl(ctx context.Context, id int) *errs.Error
}
