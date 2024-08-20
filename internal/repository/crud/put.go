package crud

import (
	"context"
	"pari/internal/domain"
)

func (r CrudRepository) PutItem(ctx context.Context, id string, updatedItem domain.Items) error {

	if err := r.db.WithContext(ctx).
		Model(&domain.Items{}).
		Where("id = ?", id).
		Updates(updatedItem).
		Error; err != nil {
		return err
	}

	return nil
}
