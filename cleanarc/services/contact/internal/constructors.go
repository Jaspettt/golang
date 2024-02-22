package internal

import (
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/usecase"
	"database/sql"
	"log/slog"
)

func NewRepository(db *sql.DB) repository.ContactRepository {
	return repository.NewRepository(db)
}

func NewUseCase(repo repository.ContactRepository) usecase.ContactUseCase {
	return usecase.NewUseCase(repo)
}

func NewDelivery(useCase usecase.ContactUseCase, logger *slog.Logger) delivery.HTTPDelivery {
	return delivery.NewDelivery(useCase, logger)
}
