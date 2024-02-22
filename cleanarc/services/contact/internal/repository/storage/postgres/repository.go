package postgres

import "architecture_go/services/contact/internal/domain/contact"

type ContactRepository interface {
	CreateContact(firstName, patronymic, phoneNumber string) (*contact.Contact, error)
	GetContactByID(id int) (*contact.Contact, error)
	UpdateContact(contact *contact.Contact) error
	DeleteContact(id int) error
}

type GroupRepository interface {
	CreateGroup(name string) (*contact.Group, error)
	GetGroupByID(id int) (*contact.Group, error)
}
