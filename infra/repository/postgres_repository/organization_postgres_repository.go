package postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/pepedoni/go-clean-arch-org-user-api/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/organization"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/infra/database/postgres"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
)

type OrganizationPostgresRepository struct {
	db postgres.PoolInterface
}

func NewOrganizationPostgresRepository(db postgres.PoolInterface) organization.OrganizationRepositoryInterface {
	return &OrganizationPostgresRepository{
		db: db,
	}
}

const (
	queryGetOrganization       = "SELECT u.id, u.name, u.document FROM organizations u WHERE u.id = $1"
	queryGetOrganizations      = "SELECT u.id, u.name, u.document FROM organizations u"
	queryGetOrganizationsCount = "SELECT COUNT(*) FROM organizations"
	queryInsertOrganization    = "INSERT INTO organizations (id, name, document) VALUES ($1, $2, $3) returning *"
	queryUpdateOrganization    = "UPDATE organizations SET name = $1, document = $2 WHERE id = $3"
	queryDeleteOrganization    = "DELETE FROM organizations WHERE id = $1"
)

func buildQueryOrganizations(filter dto.FilterGetOrganizationDTO) string {
	return queryGetOrganizations + " LIMIT " + strconv.Itoa(filter.Limit) + " OFFSET " + strconv.Itoa((filter.Page-1)*filter.Limit)
}

func (r *OrganizationPostgresRepository) Create(u *organization.Organization) (*organization.Organization, error) {
	ctx := context.Background()

	err := r.db.QueryRow(
		ctx,
		queryInsertOrganization,
		u.Id,
		u.Name,
		u.Document,
	).Scan(
		&u.Id,
		&u.Name,
		&u.Document,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *OrganizationPostgresRepository) Get(filter dto.FilterGetOrganizationDTO) (*response.PaginationReponse[[]organization.Organization], error) {
	ctx := context.Background()
	organizations := []organization.Organization{}
	total := int(0)

	query := buildQueryOrganizations(filter)

	{
		rows, err := r.db.Query(
			ctx,
			query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			organization := organization.Organization{}

			rows.Scan(
				&organization.Id,
				&organization.Name,
				&organization.Document,
			)

			organizations = append(organizations, organization)
		}
	}

	{
		err := r.db.QueryRow(ctx, queryGetOrganizationsCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &response.PaginationReponse[[]organization.Organization]{
		Items: organizations,
		Total: total,
		Page:  filter.Page,
		Limit: filter.Limit,
	}, nil
}

func (r *OrganizationPostgresRepository) GetById(id string) (*organization.Organization, error) {
	ctx := context.Background()

	u := organization.Organization{}
	fmt.Println("to aquiii", id)
	err := r.db.QueryRow(
		ctx,
		queryGetOrganization,
		id,
	).Scan(
		&u.Id,
		&u.Name,
		&u.Document,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &u, nil
}

func (r *OrganizationPostgresRepository) Delete(id string) error {
	ctx := context.Background()

	result, err := r.db.Exec(
		ctx,
		queryDeleteOrganization,
		id,
	)
	if result.RowsAffected() == 0 {
		return errors.New(constants.NOT_FOUND)
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *OrganizationPostgresRepository) Update(u *organization.Organization) (*organization.Organization, error) {
	ctx := context.Background()
	_, err := r.db.Exec(
		ctx,
		queryUpdateOrganization,
		u.Name,
		u.Document,
		u.Id,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}
