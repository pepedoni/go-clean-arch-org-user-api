package organization_user

type OrganizationUserServiceInterface interface {
	Create(organizationUser *OrganizationUser) (*OrganizationUser, error)
	Delete(userId string, orgId string) error
}

type OrganizationUserRepositoryInterface interface {
	Create(organizationUser *OrganizationUser) (*OrganizationUser, error)
	Delete(userId string, orgId string) error
}
