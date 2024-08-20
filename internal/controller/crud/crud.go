package crud

import (
	"pari/internal/domain"
)

type CrudController struct {
	crudUsecase domain.CrudUsecase
}

func New(crudUsecase domain.CrudUsecase) CrudController {
	return CrudController{
		crudUsecase: crudUsecase,
	}
}
