package domain

import (
	"VinylShop/pkg/web/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

type User struct {
	ID        int
	Email     string
	Password  []byte
	CreatedAt time.Time
}

func NewUser(email, password string) (*User, error) {
	newUser := &User{
		Email:     email,
		Password:  []byte(password),
		CreatedAt: time.Now(),
	}
	if err := newUser.IsValid(); err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(newUser.Password, 12)
	if err != nil {
		return nil, err
	}
	newUser.Password = hashedPassword
	return newUser, nil
}

var (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// passwordRegex = `^[a-zA-Z\d]*[a-z][a-zA-Z\d]*[A-Z][a-zA-Z\d]*[a-zA-Z\d]*$`
	passwordRegex = `[a-zA-Z0-9]{3,}`
)

func (u *User) SetPassword(password []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

func (u *User) IsValid() error {
	if err := u.ValidateEmail(); err != nil {
		return err
	}
	if err := u.ValidatePassword(); err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateEmail() error {
	if len(u.Email) == 0 {
		return errors.ErrEmptyField
	}
	if !regexp.MustCompile(emailRegex).MatchString(u.Email) {
		return errors.ErrEmailValidation
	}
	return nil
}

func (u *User) ValidatePassword() error {
	if len(u.Password) == 0 {
		return errors.ErrEmptyField
	}
	if !regexp.MustCompile(passwordRegex).MatchString(string(u.Password)) {
		return errors.ErrPasswordValidation
	}
	return nil
}
