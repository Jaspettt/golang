package delivery

import (
	"architecture_go/services/contact/internal/usecase"
	"log/slog"
	"net/http"
)

type HTTPDelivery struct {
	UseCase usecase.ContactUseCase
	Logger  *slog.Logger
}

func (H HTTPDelivery) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewDelivery(useCase usecase.ContactUseCase, logger *slog.Logger) HTTPDelivery {
	return HTTPDelivery{
		UseCase: useCase,
		Logger:  logger,
	}
}
