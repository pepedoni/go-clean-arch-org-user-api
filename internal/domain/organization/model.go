package organization

import "strings"

type Organization struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
}

func NewOrganization(name, document string) *Organization {
	documentAndPhoneReplacer := strings.NewReplacer("(", "", "-", "", ")", "", " ", "")
	return &Organization{
		Name:     name,
		Document: documentAndPhoneReplacer.Replace(document),
	}
}

func (o *Organization) Validate() error {
	// var err error

	// check document is valid

	return nil
}
