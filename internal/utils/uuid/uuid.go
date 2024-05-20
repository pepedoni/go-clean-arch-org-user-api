package uuid

import "github.com/google/uuid"

type UUIDGeneratorInterface interface {
	Generate() (uuid.UUID, error)
}

func New() UUIDGeneratorInterface {
	return &uuidGenerator{}
}

type uuidGenerator struct {
}

func (u *uuidGenerator) Generate() (uuid.UUID, error) {
	return uuid.NewUUID()
}
