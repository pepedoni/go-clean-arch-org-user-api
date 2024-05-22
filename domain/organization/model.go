package organization

import "strings"

type Organization struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
}

func NewOrganization(name, document string) *Organization {
	org := &Organization{
		Name:     name,
		Document: document,
	}
	org.FormatDocument()
	return org
}

func (o *Organization) FormatDocument() {
	documentReplacer := strings.NewReplacer("(", "", "-", "", ")", "", " ", "")
	o.Document = documentReplacer.Replace(o.Document)
}

func (o *Organization) Validate() error {
	return nil
}
