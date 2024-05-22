package organization_user

type OrganizationUser struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	OrganizationId string `json:"organization_id"`
}

func NewOrganizationUser(userId, organizationId string) *OrganizationUser {
	orgUser := &OrganizationUser{
		UserId:         userId,
		OrganizationId: organizationId,
	}
	return orgUser
}

func (o *OrganizationUser) Validate() error {
	return nil
}
