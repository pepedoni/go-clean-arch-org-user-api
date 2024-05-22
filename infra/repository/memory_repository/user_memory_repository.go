package memory_repository

import (
	"errors"
	"strconv"

	"github.com/pepedoni/go-clean-arch-org-user-api/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/user"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
)

type UserMemoryRepository struct{}

var usersDb = make(map[string]user.User)

func NewUserMemoryRepository() user.UserRepositoryInterface {
	return &UserMemoryRepository{}
}

func (r *UserMemoryRepository) Create(u *user.User) (*user.User, error) {
	u.Id = strconv.Itoa(len(usersDb))
	usersDb[u.Id] = *u
	return u, nil
}

func (r *UserMemoryRepository) Get(filter dto.FilterGetUserDTO) (*response.PaginationReponse[[]user.User], error) {
	users := make([]user.User, 0)

	for _, u := range usersDb {
		users = append(users, u)
	}

	resp := &response.PaginationReponse[[]user.User]{
		Items: users,
		Total: len(users),
		Page:  filter.Page,
		Limit: filter.Limit,
	}

	return resp, nil
}

func (r *UserMemoryRepository) GetById(id string) (*user.User, error) {
	u, ok := usersDb[id]
	if !ok {
		return nil, errors.New(constants.NOT_FOUND)
	}
	return &u, nil
}

func (r *UserMemoryRepository) Delete(id string) error {
	_, ok := usersDb[id]
	if !ok {
		return errors.New(constants.NOT_FOUND)
	}
	delete(usersDb, id)
	return nil
}

func (r *UserMemoryRepository) Update(u *user.User) (*user.User, error) {
	_, ok := usersDb[u.Id]
	if !ok {
		return nil, errors.New(constants.NOT_FOUND)
	}
	usersDb[u.Id] = *u
	return u, nil
}
