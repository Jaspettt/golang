package usecase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type ContactUseCase interface {
	CreateContact(firstName, patronymic, phoneNumber string) (*contact.Contact, error)
	GetContactByID(id int) (*contact.Contact, error)
	UpdateContact(contact *contact.Contact) error
	DeleteContact(id int) error
}

type GroupUseCase interface {
	CreateGroup(name string) (*group.Group, error)
	GetGroupByID(id int) (*group.Group, error)
	AddContactToGroup(contactID, groupID int) error
}

type UseCase struct {
	ContactRepo contact.ContactRepository
	GroupRepo   group.GroupRepository
}

func NewUseCase(contactRepo contact.ContactRepository, groupRepo group.GroupRepository) *UseCase {
	return &UseCase{
		ContactRepo: contactRepo,
		GroupRepo:   groupRepo,
	}
}

func (uc *UseCase) CreateContact(firstName, patronymic, phoneNumber string) (*contact.Contact, error) {

}

func (uc *UseCase) GetContactByID(id int) (*contact.Contact, error) {

}

func (uc *UseCase) UpdateContact(contact *contact.Contact) error {

}

func (uc *UseCase) DeleteContact(id int) error {

}

func (uc *UseCase) CreateGroup(name string) (*group.Group, error) {

}

func (uc *UseCase) GetGroupByID(id int) (*group.Group, error) {

}

func (uc *UseCase) AddContactToGroup(contactID, groupID int) error {

}
