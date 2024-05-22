package postgres_repository

import (
	"context"
	"errors"
	"strconv"

	"github.com/pepedoni/go-clean-arch-org-user-api/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/user"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/infra/database/postgres"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
)

type UserPostgresRepository struct {
	db postgres.PoolInterface
}

func NewUserPostgresRepository(db postgres.PoolInterface) user.UserRepositoryInterface {
	return &UserPostgresRepository{
		db: db,
	}
}

const (
	queryGetUser       = "SELECT u.id, u.name, u.email, u.phone, u.document FROM users u WHERE u.id = $1"
	queryGetUsers      = "SELECT u.id, u.name, u.email, u.phone, u.document FROM users u"
	queryGetUsersCount = "SELECT COUNT(*) FROM users"
	queryInsertUser    = "INSERT INTO users (id, name, email, phone, document) VALUES ($1, $2, $3, $4, $5) returning *"
	queryUpdateUser    = "UPDATE users SET name = $1, email = $2, phone = $3, document = $4 WHERE id = $5"
	queryDeleteUser    = "DELETE FROM users WHERE id = $1"
)

func buildQueryUsers(filter dto.FilterGetUserDTO) string {
	return queryGetUsers + " LIMIT " + strconv.Itoa(filter.Limit) + " OFFSET " + strconv.Itoa((filter.Page-1)*filter.Limit)
}

func (r *UserPostgresRepository) Create(u *user.User) (*user.User, error) {
	ctx := context.Background()

	err := r.db.QueryRow(
		ctx,
		queryInsertUser,
		u.Id,
		u.Name,
		u.Email,
		u.Phone,
		u.Document,
	).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Phone,
		&u.Document,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserPostgresRepository) Get(filter dto.FilterGetUserDTO) (*response.PaginationReponse[[]user.User], error) {
	ctx := context.Background()
	users := []user.User{}
	total := int(0)

	query := buildQueryUsers(filter)

	{
		rows, err := r.db.Query(
			ctx,
			query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			user := user.User{}

			rows.Scan(
				&user.Id,
				&user.Name,
				&user.Email,
				&user.Phone,
				&user.Document,
			)

			users = append(users, user)
		}
	}

	{
		err := r.db.QueryRow(ctx, queryGetUsersCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &response.PaginationReponse[[]user.User]{
		Items: users,
		Total: total,
		Page:  filter.Page,
		Limit: filter.Limit,
	}, nil
}

func (r *UserPostgresRepository) GetById(id string) (*user.User, error) {
	ctx := context.Background()

	u := user.User{}

	err := r.db.QueryRow(
		ctx,
		queryGetUser,
		id,
	).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Phone,
		&u.Document,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserPostgresRepository) Delete(id string) error {
	ctx := context.Background()

	result, err := r.db.Exec(
		ctx,
		queryDeleteUser,
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

func (r *UserPostgresRepository) Update(u *user.User) (*user.User, error) {
	ctx := context.Background()
	_, err := r.db.Exec(
		ctx,
		queryUpdateUser,
		u.Name,
		u.Email,
		u.Phone,
		u.Document,
		u.Id,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}
