package postgres_repository

import (
	"context"
	"errors"

	"github.com/pepedoni/go-clean-arch-org-user-api/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/organization_user"
	"github.com/pepedoni/go-clean-arch-org-user-api/infra/database/postgres"
)

type OrganizationUserPostgresRepository struct {
	db postgres.PoolInterface
}

func NewOrganizationUserPostgresRepository(db postgres.PoolInterface) organization_user.OrganizationUserRepositoryInterface {
	return &OrganizationUserPostgresRepository{
		db: db,
	}
}

const (
	queryInsertOrganizationUser = "INSERT INTO organizations_users (id, user_id, organization_id) VALUES ($1, $2, $3) returning *"
	queryDeleteOrganizationUser = "DELETE FROM organizations_users WHERE user_id = $1 and organization_id = $2"
)

func (r *OrganizationUserPostgresRepository) Create(organizationUserCreate *organization_user.OrganizationUser) (*organization_user.OrganizationUser, error) {
	ctx := context.Background()

	err := r.db.QueryRow(
		ctx,
		queryInsertOrganizationUser,
		organizationUserCreate.Id,
		organizationUserCreate.UserId,
		organizationUserCreate.OrganizationId,
	).Scan(
		&organizationUserCreate.Id,
		&organizationUserCreate.UserId,
		&organizationUserCreate.OrganizationId,
	)

	if err != nil {
		return nil, err
	}

	return organizationUserCreate, nil
}

func (r *OrganizationUserPostgresRepository) Delete(userId, orgId string) error {
	ctx := context.Background()

	result, err := r.db.Exec(
		ctx,
		queryDeleteOrganizationUser,
		userId,
		orgId,
	)

	if result.RowsAffected() == 0 {
		return errors.New(constants.NOT_FOUND)
	}

	if err != nil {
		return err
	}

	return nil
}
