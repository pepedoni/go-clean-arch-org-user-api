package memory_repository

import (
	"errors"
	"strconv"

	"github.com/pepedoni/go-clean-arch-org-user-api/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/organization"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
)

type OrganizationMemoryRepository struct{}

var organizationsDb = make(map[string]organization.Organization)
var length = 0

func NewOrganizationMemoryRepository() organization.OrganizationRepositoryInterface {
	return &OrganizationMemoryRepository{}
}

func (r *OrganizationMemoryRepository) Create(u *organization.Organization) (*organization.Organization, error) {
	u.Id = strconv.Itoa(len(organizationsDb))
	organizationsDb[u.Id] = *u
	return u, nil
}

func (r *OrganizationMemoryRepository) Get(filter dto.FilterGetOrganizationDTO) (*response.PaginationReponse[[]organization.Organization], error) {
	organizations := make([]organization.Organization, 0)

	for _, u := range organizationsDb {
		organizations = append(organizations, u)
	}

	organizations = organizations[filter.Limit*(filter.Page-1) : filter.Limit*filter.Page]

	resp := &response.PaginationReponse[[]organization.Organization]{
		Items: organizations,
		Total: len(organizations),
		Page:  filter.Page,
		Limit: filter.Limit,
	}

	return resp, nil
}

func (r *OrganizationMemoryRepository) GetById(id string) (*organization.Organization, error) {
	u, ok := organizationsDb[id]
	if !ok {
		return nil, errors.New(constants.NOT_FOUND)
	}
	return &u, nil
}

func (r *OrganizationMemoryRepository) Delete(id string) error {
	_, ok := organizationsDb[id]
	if !ok {
		return errors.New(constants.NOT_FOUND)
	}
	delete(organizationsDb, id)
	return nil
}

func (r *OrganizationMemoryRepository) Update(u *organization.Organization) (*organization.Organization, error) {
	_, ok := organizationsDb[u.Id]
	if !ok {
		return nil, errors.New(constants.NOT_FOUND)
	}
	organizationsDb[u.Id] = *u
	return u, nil
}
