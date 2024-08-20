package crud

import (
	"context"
	"pari/internal/domain"
)

func (u CrudUsecase) GetAll(ctx context.Context) ([]domain.Items, error) {

	result, err := u.crudRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (u CrudUsecase) GetById(ctx context.Context, id string) (domain.Items, error) {

	result, err := u.crudRepo.GetById(ctx, id)
	if err != nil {
		return domain.Items{}, err
	}

	return result, nil
}
