package crud

import (
	"context"
)

func (uc CrudUsecase) DeleteItem(ctx context.Context, id string) error {
	if err := uc.crudRepo.DeleteItem(ctx, id); err != nil {
		return err
	}

	return nil
}
