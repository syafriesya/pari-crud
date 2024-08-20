package crud

import (
	"context"
	"pari/internal/domain"
)

func (r CrudRepository) PostItem(ctx context.Context, newItem domain.Items) error {
	if err := r.db.WithContext(ctx).Create(&newItem).Error; err != nil {
		return err
	}
	return nil
}
