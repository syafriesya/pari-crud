package crud

import (
	"context"
	"pari/internal/domain"
)

func (r CrudRepository) DeleteItem(ctx context.Context, id string) error {
	// Delete the item by ID
	if err := r.db.WithContext(ctx).Delete(&domain.Items{}, id).Error; err != nil {
		return err
	}

	return nil
}
