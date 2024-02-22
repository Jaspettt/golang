package usecase

import (
	"architecture_go/services/contact/internal/domain/contact"
)

type GroupUseCase interface {
	CreateGroup(name string) (*contact.Group, error)
	GetGroupByID(id int) (*contact.Group, error)
}
