package memory_repository

import (
	"errors"
	"strconv"

	"github.com/pepedoni/go-clean-arch-org-user-api/internal/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/domain/organization"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/response"
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

func (r *OrganizationMemoryRepository) Get(page int, limit int) (*response.PaginationReponse[[]organization.Organization], error) {
	organizations := make([]organization.Organization, 0)

	for _, u := range organizationsDb {
		organizations = append(organizations, u)
	}

	resp := &response.PaginationReponse[[]organization.Organization]{
		Items: organizations,
		Total: len(organizations),
		Page:  page,
		Limit: limit,
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
