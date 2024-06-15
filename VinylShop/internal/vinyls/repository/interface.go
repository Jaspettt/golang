package repository

import "VinylShop/internal/vinyls/domain"

type VinylsRepository interface {
	Insert(vinylInsert *domain.Vinyl) (int, error)
	Get(id int) (domain.Vinyl, error)
	GetAll(filters *domain.Filters) ([]domain.Vinyl, error)
	Length() int
	//Add(vinyl domain.Vinyl) (int64, error)
	Update(id int, vinylUpdate *domain.Vinyl) (*domain.Vinyl, error)
	Delete(id int) error
	Exists(id int) bool
}
