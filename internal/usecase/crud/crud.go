package crud

import (
	"pari/internal/domain"
)

type CrudUsecase struct {
	crudRepo domain.CrudRepository
}

func New(
	crudRepo domain.CrudRepository,
) CrudUsecase {
	return CrudUsecase{
		crudRepo: crudRepo,
	}
}
