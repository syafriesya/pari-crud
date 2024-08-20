package crud

import (
	"context"
	"pari/internal/domain"
)

func (r CrudRepository) PutItem(ctx context.Context, updatedItem domain.Items) error {
	var item domain.Items

	if err := r.db.WithContext(ctx).Model(&item).Updates(updatedItem).Where("id = ?", updatedItem.ID).Error; err != nil {
		return err
	}

	return nil
}
