package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/numeric"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Document string `json:"document" binding:"required"`
}

func NewUser(name, email, phone, document string) *User {
	documentAndPhoneReplacer := strings.NewReplacer("(", "", "-", "", ")", "", " ", "")
	fmt.Println(documentAndPhoneReplacer.Replace(phone))
	return &User{
		Name:     name,
		Email:    email,
		Phone:    documentAndPhoneReplacer.Replace(phone),
		Document: documentAndPhoneReplacer.Replace(document),
	}
}

func (u *User) Validate() error {
	var err error

	if !numeric.IsNumeric(u.Phone) {
		err = errors.New("invalid phone")
		return err
	}

	return nil
}
