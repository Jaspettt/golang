package delivery

import (
	"VinylShop/internal/vinyls/usecase"
	"database/sql"
)

type VinylsDeliveryImpl struct {
	usecase usecase.VinylsUseCase
}

func NewVinylsDelivery(db *sql.DB) *VinylsDeliveryImpl {
	return &VinylsDeliveryImpl{
		usecase: usecase.NewVinylUseCase(db),
	}
}
