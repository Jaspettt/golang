package usecase

import (
	"VinylShop/internal/vinyls/domain"
	repo "VinylShop/internal/vinyls/repository"
	ctx "VinylShop/pkg/types/context"
	"VinylShop/pkg/types/responses"
	errs "VinylShop/pkg/web/errors"
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type VinylUseCaseImpl struct {
	repo repo.VinylsRepository
}

func NewVinylUseCase(db *sql.DB) *VinylUseCaseImpl {
	return &VinylUseCaseImpl{
		repo: repo.NewVinylRepository(db),
	}
}
func (v *VinylUseCaseImpl) InsertVinyl(ctx context.Context, vinyl *domain.Vinyl) (int, *errs.Error) {

	if err := vinyl.Validate(); err != nil {
		return 0, errs.ErrUnpocessableEntity
	}
	id, error := v.repo.Insert(vinyl)
	if error != nil {
		return 0, errs.ErrInternal
	}
	return id, nil
}
func (v *VinylUseCaseImpl) GetVinyl(ctx context.Context, id int) (*domain.Vinyl, *errs.Error) {
	if !v.repo.Exists(id) {
		return nil, errs.ErrNotFound
	}
	vinyl, err := v.repo.Get(id)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("Problem with sql query")
		return nil, errs.ErrInternal
	}
	return &vinyl, nil
}
func (v *VinylUseCaseImpl) GetVinyls(ctx ctx.Context, filters *domain.Filters) (*responses.Pagination, *errs.Error) {
	vinyls, err := v.repo.GetAll(filters)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("Problem with sql query")
		return nil, errs.ErrInternal
	}
	response := &responses.Pagination{}
	response.Content = vinyls
	if filters.Offset < v.repo.Length() {
		response.Next = fmt.Sprintf("%s?limit=%d&offset=%d", ctx.Data["url"], filters.Limit, filters.Limit+filters.Offset)
	}
	if filters.Offset >= filters.Limit {
		response.Prev = fmt.Sprintf("%s?limit=%d&offset=%d", ctx.Data["url"], filters.Limit, filters.Offset-filters.Limit)
	} else if filters.Offset > 0 {
		response.Prev = fmt.Sprintf("%s?limit=%d&offset=%d", ctx.Data["url"], filters.Limit, 0)
	}
	return response, nil
}
func (v *VinylUseCaseImpl) UpdateVinyl(ctx context.Context, id int, vinyl *domain.Vinyl) (*domain.Vinyl, *errs.Error) {
	if !v.repo.Exists(id) {
		return nil, errs.ErrNotFound
	}
	updatedVinyl, err := v.repo.Update(id, vinyl)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("Problem with sql query")
		return nil, errs.ErrInternal
	}
	return updatedVinyl, nil
}
func (v *VinylUseCaseImpl) DeleteVinyl(ctx context.Context, id int) *errs.Error {
	if id < 0 {
		return errs.ErrBadRequest
	}
	if !v.repo.Exists(id) {
		return errs.ErrNotFound
	}
	if err := v.repo.Delete(id); err != nil {
		logrus.WithField("err", err.Error()).Error("Problem with sql query")
		return errs.ErrInternal
	}
	return nil
}
