package crud

import (
	"context"
	"errors"
	"pari/internal/domain"
)

func (uc CrudUsecase) PutItem(ctx context.Context, id string, updatedItem domain.Items) error {
	if updatedItem.Detail == "" {
		return errors.New("item detail cannot be empty")
	}

	err := uc.crudRepo.PutItem(ctx, id, updatedItem)
	if err != nil {
		return err
	}

	return nil
}
