package crud

import (
	"context"
	"errors"
	"pari/internal/domain"
)

func (uc *CrudUsecase) PostItem(ctx context.Context, newItem domain.Items) error {
	if newItem.Detail == "" {
		return errors.New("item detail cannot be empty")
	}

	err := uc.crudRepo.PostItem(ctx, newItem)
	if err != nil {
		return err
	}

	return nil
}
