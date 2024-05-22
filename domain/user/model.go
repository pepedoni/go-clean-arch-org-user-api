package user

import (
	"errors"
	"strings"

	"github.com/pepedoni/go-clean-arch-org-user-api/utils/numeric"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Document string `json:"document" binding:"required"`
}

func NewUser(name, email, phone, document string) *User {
	newUser := &User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Document: document,
	}
	newUser.FormatPhone()
	newUser.FormatDocument()
	return newUser
}

func (u *User) FormatPhone() {
	phoneReplacer := strings.NewReplacer("(", "", "-", "", ")", "", " ", "")
	u.Phone = phoneReplacer.Replace(u.Phone)
}

func (u *User) FormatDocument() {
	documentReplacer := strings.NewReplacer("(", "", "-", "", ")", "", " ", "")
	u.Phone = documentReplacer.Replace(u.Phone)
}

func (u *User) Validate() error {
	var err error

	if !numeric.IsNumeric(u.Phone) {
		err = errors.New("invalid phone")
		return err
	}

	return nil
}
