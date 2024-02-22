package contact

import (
	"fmt"
	"strings"
)

type Contact struct {
	ID          int
	FirstName   string
	Patronymic  string
	PhoneNumber string
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.Patronymic)
}

func (c *Contact) SetFullName(fullName string) {
	parts := strings.Fields(fullName)
	if len(parts) >= 2 {
		c.FirstName = parts[0]
		c.Patronymic = parts[1]
	}
}

func (c *Contact) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = phoneNumber
}

type Group struct {
	ID   int
	Name string
}
